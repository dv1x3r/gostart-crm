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

type Attribute struct {
	attributeGroupService *service.AttributeGroup
	attributeSetService   *service.AttributeSet
	attributeValueService *service.AttributeValue
}

func NewAttribute(
	attributeGroupService *service.AttributeGroup,
	attributeSetService *service.AttributeSet,
	attributeValueService *service.AttributeValue,
) *Attribute {
	return &Attribute{
		attributeGroupService: attributeGroupService,
		attributeSetService:   attributeSetService,
		attributeValueService: attributeValueService,
	}
}

func (ep *Attribute) Register(g *echo.Group) {
	g.GET("/group", ep.GetGroup)
	g.GET("/group/dropdown", ep.GetGroupDropdown)
	g.POST("/group/save", ep.PostGroupSave)
	g.POST("/group/:groupID/delete", ep.PostGroupDelete)
	g.GET("/group/:groupID/set", ep.GetSet)
	g.POST("/group/:groupID/set/save", ep.PostSetSave)
	g.POST("/set/delete", ep.PostSetDelete)
	g.POST("/set/reorder", ep.PostSetReorder)
	g.GET("/set/:setID/value", ep.GetValue)
	g.GET("/set/:setID/value/dropdown", ep.GetValueDropdown)
	g.POST("/set/:setID/value/save", ep.PostValueSave)
	g.POST("/value/delete", ep.PostValueDelete)
	g.POST("/value/reorder", ep.PostValueReorder)
}

func (ep *Attribute) GetGroup(c echo.Context) error {
	rows, err := ep.attributeGroupService.FetchAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.AttributeGroupW2GridResponse{Status: "success", Records: rows})
}

func (ep *Attribute) GetGroupDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.attributeGroupService.FetchDropdown(c.Request().Context(), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownBasicResponse{Status: "success", Records: rows})
}

func (ep *Attribute) PostGroupSave(c echo.Context) error {
	dto := &model.AttributeGroup{}
	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	err := ep.attributeGroupService.Save(c.Request().Context(), *dto)
	if errors.Is(err, service.ErrAttributeGroupExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The group with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) PostGroupDelete(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err = ep.attributeGroupService.Delete(c.Request().Context(), int64(groupID))
	if errors.Is(err, service.ErrAttributeGroupHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The group contains related objects, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) GetSet(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.attributeSetService.Fetch(c.Request().Context(), int64(groupID), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.AttributeSetW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Attribute) PostSetSave(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.AttributeSetW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err = ep.attributeSetService.SavePartials(c.Request().Context(), int64(groupID), req.Changes)
	if errors.Is(err, service.ErrAttributeSetExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The set with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) PostSetDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.attributeSetService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrAttributeSetHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The set contains related objects, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) PostSetReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.attributeSetService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrAttributeSetNotFound) {
		return c.JSON(http.StatusNotFound, model.ResponseErrorJson("Set not found"))
	} else if errors.Is(err, service.ErrAttributeSetInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) GetValue(c echo.Context) error {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.attributeValueService.Fetch(c.Request().Context(), int64(setID), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.AttributeValueW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Attribute) GetValueDropdown(c echo.Context) error {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.attributeValueService.FetchDropdownBySetID(c.Request().Context(), int64(setID), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownBasicResponse{Status: "success", Records: rows})
}

func (ep *Attribute) PostValueSave(c echo.Context) error {
	setID, err := strconv.Atoi(c.Param("setID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.AttributeValueW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err = ep.attributeValueService.SavePartials(c.Request().Context(), int64(setID), req.Changes)
	if errors.Is(err, service.ErrAttributeValueExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The value with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) PostValueDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.attributeValueService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrAttributeValueHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("The attribute value contains related objects, please unlink them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Attribute) PostValueReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.attributeValueService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrAttributeValueNotFound) {
		return c.JSON(http.StatusNotFound, model.ResponseErrorJson("Value not found"))
	} else if errors.Is(err, service.ErrAttributeValueInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}
