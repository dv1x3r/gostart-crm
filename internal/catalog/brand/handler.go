package brand

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
	g.GET("/dropdown", h.GetDropdown)
	g.POST("/save", h.PostSave)
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

func (h *Handler) GetDropdown(c echo.Context) error {
	req, err := w2.ParseDropdownRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, err := h.service.FetchDropdown(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewDropdownResponse(rows))
}

func (h *Handler) PostSave(c echo.Context) error {
	req, err := w2.ParseGridSaveRequest[Brand](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.SaveChanges(c.Request().Context(), req.Changes)
	if errors.Is(err, ErrBrandExists) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The brand with the same name already exists"))
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
	if errors.Is(err, ErrBrandHasRelated) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("Brand contains related products, please remove them first"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}
