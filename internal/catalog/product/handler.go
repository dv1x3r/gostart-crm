package product

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
	g.GET("", h.Get)
	g.POST("", h.Post)
	g.DELETE("/:productID", h.Delete)
	g.GET("/category/:categoryID/catalog", h.GetCategoryCatalog)
	g.POST("/save", h.PostSave)
	g.POST("/delete", h.PostDelete)
	g.GET("/:productID/attributes", h.GetAttributes)
	g.POST("/:productID/attributes", h.PostAttributes)
}

func (h *Handler) Get(c echo.Context) error {
	req, err := w2.ParseFormGetRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	product, err := h.service.FetchByID(c.Request().Context(), req.RecID)
	if errors.Is(err, ErrProductNotFound) {
		return c.JSON(http.StatusNotFound, w2.NewErrorResponse("Product not found"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewFormGetResponse(product))
}

func (h *Handler) Post(c echo.Context) error {
	req, err := w2.ParseFormSaveRequest[Product](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	id, err := h.service.SaveForm(c.Request().Context(), req.Record)
	if errors.Is(err, ErrProductExists) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The product with the same supplier code already exists"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewFormSaveResponse(id))
}

func (h *Handler) Delete(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.DeleteListID(c.Request().Context(), []int{productID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) GetCategoryCatalog(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridDataRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, count, err := h.service.Fetch(c.Request().Context(), req, categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewGridDataResponse(rows, count))
}

func (h *Handler) PostSave(c echo.Context) error {
	req, err := w2.ParseGridSaveRequest[Product](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.SaveChanges(c.Request().Context(), req.Changes)
	if err != nil {
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
	if errors.Is(err, ErrProductDeleteProtect) {
		return c.JSON(http.StatusUnprocessableEntity, w2.NewErrorResponse("You can delete only up to 50 products per request"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) GetAttributes(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridDataRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, count, err := h.service.FetchAttributes(c.Request().Context(), productID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewGridDataResponse(rows, count))
}

func (h *Handler) PostAttributes(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridSaveRequest[ProductAttribute](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.SaveAttributeChanges(c.Request().Context(), productID, req.Changes)
	if errors.Is(err, ErrProductAttributeNotFound) {
		return c.JSON(http.StatusNotFound, w2.NewErrorResponse("Product attribute not found"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}
