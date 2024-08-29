package endpoint

import (
	"encoding/json"
	"errors"
	"net/http"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/service"

	"github.com/labstack/echo/v4"
)

type Supplier struct {
	supplierService *service.Supplier
}

func NewSupplier(supplierService *service.Supplier) *Supplier {
	return &Supplier{supplierService: supplierService}
}

func (ep *Supplier) Register(g *echo.Group) {
	g.GET("", ep.Get)
	g.GET("/dropdown", ep.GetDropdown)
	g.POST("/save", ep.PostSave)
	g.POST("/delete", ep.PostDelete)
	g.POST("/reorder", ep.PostReorder)
}

func (ep *Supplier) Get(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.supplierService.Fetch(c.Request().Context(), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.SupplierW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Supplier) GetDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.supplierService.FetchDropdown(c.Request().Context(), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownBasicResponse{Status: "success", Records: rows})
}

func (ep *Supplier) PostSave(c echo.Context) error {
	req := &model.SupplierW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err := ep.supplierService.SavePartials(c.Request().Context(), req.Changes)
	if errors.Is(err, service.ErrSupplierExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The supplier with the same name or code already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Supplier) PostDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.supplierService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrSupplierHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Supplier contains related products, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Supplier) PostReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.supplierService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrSupplierInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}
