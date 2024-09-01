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

type Product struct {
	productService       *service.Product
	productStatusService *service.ProductStatus
}

func NewProduct(productService *service.Product, productStatusService *service.ProductStatus) *Product {
	return &Product{
		productService:       productService,
		productStatusService: productStatusService,
	}
}

func (ep *Product) Register(g *echo.Group) {
	g.GET("", ep.Get)
	g.POST("", ep.Post)
	g.DELETE("/:productID", ep.Delete)
	g.GET("/category/:categoryID/catalog", ep.GetCategoryCatalog)
	g.POST("/save", ep.PostSave)
	g.POST("/delete", ep.PostDelete)
	g.GET("/:productID/attributes", ep.GetAttributes)
	g.POST("/:productID/attributes", ep.PostAttributes)
	g.GET("/status", ep.GetStatus)
	g.GET("/status/dropdown", ep.GetStatusDropdown)
	g.POST("/status/save", ep.PostStatusSave)
	g.POST("/status/delete", ep.PostStatusDelete)
	g.POST("/status/reorder", ep.PostStatusReorder)
}

func (ep *Product) Get(c echo.Context) error {
	req := &model.ProductW2FormRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	product, err := ep.productService.FetchByID(c.Request().Context(), req.RecID)
	if errors.Is(err, service.ErrProductNotFound) {
		return c.JSON(http.StatusNotFound, model.ResponseErrorJson("Product not found"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ProductW2FormResponse{Status: "success", Record: &product})
}

func (ep *Product) Post(c echo.Context) error {
	req := &model.ProductW2FormRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	req.Record.ID = req.RecID

	id, err := ep.productService.SaveForm(c.Request().Context(), req.Record)
	if errors.Is(err, service.ErrProductExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The product with the same supplier code already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ProductW2FormResponse{Status: "success", RecID: id})
}

func (ep *Product) Delete(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	if _, err = ep.productService.DeleteListID(c.Request().Context(), []int64{int64(productID)}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Product) GetCategoryCatalog(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.productService.Fetch(c.Request().Context(), req.ToFindManyParams(), int64(categoryID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ProductW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Product) PostSave(c echo.Context) error {
	req := &model.ProductW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	if _, err := ep.productService.SavePartials(c.Request().Context(), req.Changes); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Product) PostDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.productService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrProductDeleteProtect) {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseErrorJson("You can delete only up to 50 products per request"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Product) GetAttributes(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.productService.FetchAttributes(c.Request().Context(), int64(productID), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ProductAttributeW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Product) PostAttributes(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.ProductAttributeW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err = ep.productService.SaveAttributePartials(c.Request().Context(), int64(productID), req.Changes)
	if errors.Is(err, service.ErrProductAttributeNotFound) {
		return c.JSON(http.StatusNotFound, model.ResponseErrorJson("Product attribute not found"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Product) GetStatus(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.productStatusService.Fetch(c.Request().Context(), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ProductStatusW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Product) GetStatusDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.productStatusService.FetchDropdown(c.Request().Context(), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownGenericResponse[model.ProductStatusDropdown]{Status: "success", Records: rows})
}

func (ep *Product) PostStatusSave(c echo.Context) error {
	req := &model.ProductStatusW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err := ep.productStatusService.SavePartials(c.Request().Context(), req.Changes)
	if errors.Is(err, service.ErrProductStatusExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Status with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Product) PostStatusDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.productStatusService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrProductStatusHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Status contains related products, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Product) PostStatusReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.productStatusService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrProductStatusInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}
