package attribute

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
	g.GET("/group", h.GetGroup)
	g.GET("/group/dropdown", h.GetGroupDropdown)
	g.POST("/group/save", h.PostGroupSave)
	g.POST("/group/:groupID/delete", h.PostGroupDelete)
	g.GET("/group/:groupID/set", h.GetSet)
	g.POST("/group/:groupID/set/save", h.PostSetSave)
	g.POST("/set/delete", h.PostSetDelete)
	g.POST("/set/reorder", h.PostSetReorder)
	g.GET("/set/:setID/value", h.GetValue)
	g.GET("/set/:setID/value/dropdown", h.GetValueDropdown)
	g.POST("/set/:setID/value/save", h.PostValueSave)
	g.POST("/value/delete", h.PostValueDelete)
	g.POST("/value/reorder", h.PostValueReorder)
}

func (h *Handler) GetGroup(c echo.Context) error {
	rows, err := h.service.FetchAllGroups(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, w2.NewDropdownResponse(rows))
}

func (h *Handler) GetGroupDropdown(c echo.Context) error {
	req, err := w2.ParseDropdownRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, err := h.service.FetchGroupDropdown(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewDropdownResponse(rows))
}

func (h *Handler) PostGroupSave(c echo.Context) error {
	group := AttributeGroup{}
	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err := h.service.SaveGroup(c.Request().Context(), group)
	if errors.Is(err, ErrAttributeExists) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The group with the same name already exists"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostGroupDelete(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.DeleteGroup(c.Request().Context(), groupID)
	if errors.Is(err, ErrAttributeHasRelated) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The group contains related objects, please remove them first"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) GetSet(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridDataRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, count, err := h.service.FetchSets(c.Request().Context(), groupID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewGridDataResponse(rows, count))
}

func (h *Handler) PostSetSave(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridSaveRequest[AttributeSet](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.SaveSetChanges(c.Request().Context(), groupID, req.Changes)
	if errors.Is(err, ErrAttributeExists) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The set with the same name already exists"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostSetDelete(c echo.Context) error {
	req, err := w2.ParseGridRemoveRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.DeleteSetListID(c.Request().Context(), req.ID)
	if errors.Is(err, ErrAttributeHasRelated) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The set contains related objects, please remove them first"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostSetReorder(c echo.Context) error {
	req, err := w2.ParseGridReorderRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.ReorderSets(c.Request().Context(), req)
	if errors.Is(err, ErrAttributeNotFound) {
		return c.JSON(http.StatusNotFound, w2.NewErrorResponse("Set not found"))
	} else if errors.Is(err, ErrAttributeInvalidReorder) {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse("Invalid reorder parameters"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) GetValue(c echo.Context) error {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridDataRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, count, err := h.service.FetchValues(c.Request().Context(), setID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewGridDataResponse(rows, count))
}

func (h *Handler) GetValueDropdown(c echo.Context) error {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseDropdownRequest(c.QueryParam("request"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	rows, err := h.service.FetchValueDropdownBySetID(c.Request().Context(), setID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewDropdownResponse(rows))
}

func (h *Handler) PostValueSave(c echo.Context) error {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	req, err := w2.ParseGridSaveRequest[AttributeValue](c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.SaveValueChanges(c.Request().Context(), setID, req.Changes)
	if errors.Is(err, ErrAttributeExists) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The value with the same name already exists"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostValueDelete(c echo.Context) error {
	req, err := w2.ParseGridRemoveRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.DeleteValueListID(c.Request().Context(), req.ID)
	if errors.Is(err, ErrAttributeHasRelated) {
		return c.JSON(http.StatusConflict, w2.NewErrorResponse("The attribute value contains related objects, please unlink them first"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}

func (h *Handler) PostValueReorder(c echo.Context) error {
	req, err := w2.ParseGridReorderRequest(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse(err.Error()))
	}

	err = h.service.ReorderValues(c.Request().Context(), req)
	if errors.Is(err, ErrAttributeNotFound) {
		return c.JSON(http.StatusNotFound, w2.NewErrorResponse("Value not found"))
	} else if errors.Is(err, ErrAttributeInvalidReorder) {
		return c.JSON(http.StatusBadRequest, w2.NewErrorResponse("Invalid reorder parameters"))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, w2.NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, w2.NewSuccessResponse())
}
