package endpoint

import (
	"encoding/json"
	"errors"
	"net/http"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/service"

	"github.com/labstack/echo/v4"
)

type Brand struct {
	brandService *service.Brand
}

func NewBrand(brandService *service.Brand) *Brand {
	return &Brand{brandService: brandService}
}

func (ep *Brand) Register(g *echo.Group) {
	g.GET("", ep.Get)
	g.GET("/dropdown", ep.GetDropdown)
	g.POST("/save", ep.PostSave)
	g.POST("/delete", ep.PostDelete)
}

func (ep *Brand) Get(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.brandService.Fetch(c.Request().Context(), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.BrandW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Brand) GetDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.brandService.FetchDropdown(c.Request().Context(), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownBasicResponse{Status: "success", Records: rows})
}

func (ep *Brand) PostSave(c echo.Context) error {
	req := &model.BrandW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err := ep.brandService.SavePartials(c.Request().Context(), req.Changes)
	if errors.Is(err, service.ErrBrandExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The brand with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Brand) PostDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.brandService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrBrandHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Brand contains related products, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}
