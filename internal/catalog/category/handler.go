package category

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dv1x3r/w2go/w2"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(g *echo.Group) {
	g.GET("/tree", h.GetTree)
	g.GET("/dropdown", h.GetDropdown)
	g.GET("/:parentID/children", h.GetChildren)
	g.POST("/save", h.PostSave)
	g.POST("/delete", h.PostDelete)
	g.POST("/reorder", h.PostReorder)
}

func (h *Handler) GetTree(c echo.Context) error {
	tree, err := h.service.FetchTree(c.Request().Context(), false)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]any{"status": w2.StatusSuccess, "tree": tree})
}

func (h *Handler) GetDropdown(c echo.Context) error {
	req, err := w2.ParseDropdownRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, err := h.service.FetchDropdown(c.Request().Context(), req, c.QueryParam("leafs") == "1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewDropdownResponse(rows))
}

func (h *Handler) GetChildren(c echo.Context) error {
	parentID, err := strconv.Atoi(c.Param("parentID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridDataRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, count, err := h.service.FetchChildren(c.Request().Context(), parentID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewGridDataResponse(rows, count))
}

func (h *Handler) PostSave(c echo.Context) error {
	req, err := w2.ParseGridSaveRequest[Category](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.SaveChanges(c.Request().Context(), req.Changes)
	if errors.Is(err, ErrCategoryExists) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The category with the same name already exists"))
	} else if errors.Is(err, ErrCategoryInvalidCircular) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The category cannot have itself as a parent"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostDelete(c echo.Context) error {
	req, err := w2.ParseGridRemoveRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, ErrCategoryHasRelated) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The category contains related objects, please remove them first"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostReorder(c echo.Context) error {
	req, err := w2.ParseGridReorderRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.Reorder(c.Request().Context(), req)
	if errors.Is(err, ErrCategoryNotFound) {
		return c.JSON(http.StatusNotFound, w2.NewErrorResponse("Category not found"))
	} else if errors.Is(err, ErrCategoryInvalidReorder) {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse("Invalid reorder parameters"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}
