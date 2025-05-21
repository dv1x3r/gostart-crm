package order

import (
	"errors"
	"net/http"

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
	g.POST("/delete", h.PostDelete)
}

func (h *Handler) Get(c echo.Context) error {
	req, err := w2.ParseGridDataRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, count, err := h.service.Fetch(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewGridDataResponse(rows, count))
}

func (h *Handler) Post(c echo.Context) error {
	req, err := w2.ParseFormSaveRequest[Order](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	id, err := h.service.SaveForm(c.Request().Context(), req.RecID, req.Record)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewFormSaveResponse(id))
}

func (h *Handler) PostDelete(c echo.Context) error {
	req, err := w2.ParseGridRemoveRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, ErrOrderDeleteProtect) {
		return c.JSON(http.StatusUnprocessableEntity, w2.NewErrorResponse("You can delete only 20 orders per request"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}
