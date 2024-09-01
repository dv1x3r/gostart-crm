package endpoint

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/service"

	"github.com/labstack/echo/v4"
)

type Category struct {
	categoryService *service.Category
}

func NewCategory(categoryService *service.Category) *Category {
	return &Category{categoryService: categoryService}
}

func (ep *Category) Register(g *echo.Group) {
	g.GET("/tree", ep.GetTree)
	g.GET("/dropdown", ep.GetDropdown)
	g.GET("/:parentID/children", ep.GetChildren)
	g.POST("/save", ep.PostSave)
	g.POST("/delete", ep.PostDelete)
	g.POST("/reorder", ep.PostReorder)
}

func (ep *Category) GetTree(c echo.Context) error {
	categoryTree, err := ep.categoryService.FetchTree(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(categoryTree))
}

func (ep *Category) GetDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.categoryService.FetchDropdown(c.Request().Context(), req.Search, req.Max, c.QueryParam("leafs") == "1")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownGenericResponse[model.CategoryDropdown]{Status: "success", Records: rows})
}

func (ep *Category) GetChildren(c echo.Context) error {
	parentID, err := strconv.Atoi(c.Param("parentID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.categoryService.FetchChildren(c.Request().Context(), int64(parentID), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.CategoryW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Category) PostSave(c echo.Context) error {
	req := &model.CategoryW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err := ep.categoryService.SavePartials(c.Request().Context(), req.Changes)
	if errors.Is(err, service.ErrCategoryExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The category with the same name already exists"))
	} else if errors.Is(err, service.ErrCategoryInvalidCircular) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The category cannot have itself as a parent"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Category) PostDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.categoryService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrCategoryHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The category contains related objects, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Category) PostReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.categoryService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrCategoryNotFound) {
		return c.JSON(http.StatusNotFound, model.ResponseErrorJson("Category not found"))
	} else if errors.Is(err, service.ErrCategoryInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}
