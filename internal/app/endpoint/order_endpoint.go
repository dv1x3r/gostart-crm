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

type Order struct {
	orderService         *service.Order
	orderStatusService   *service.OrderStatus
	paymentMethodService *service.PaymentMethod
}

func NewOrder(
	orderService *service.Order,
	orderStatusService *service.OrderStatus,
	paymentMethodService *service.PaymentMethod,
) *Order {
	return &Order{
		orderService:         orderService,
		orderStatusService:   orderStatusService,
		paymentMethodService: paymentMethodService,
	}
}

func (ep *Order) Register(g *echo.Group) {
	g.GET("", ep.Get)
	g.POST("", ep.Post)
	g.POST("/delete", ep.PostDelete)
	g.GET("/counter", ep.GetCounter)
	g.GET("/:orderID/line", ep.GetLine)
	g.GET("/status", ep.GetStatus)
	g.GET("/status/dropdown", ep.GetStatusDropdown)
	g.POST("/status/save", ep.PostStatusSave)
	g.POST("/status/delete", ep.PostStatusDelete)
	g.POST("/status/reorder", ep.PostStatusReorder)
	g.GET("/payment", ep.GetPayment)
	g.GET("/payment/dropdown", ep.GetPaymentDropdown)
	g.POST("/payment/save", ep.PostPaymentSave)
	g.POST("/payment/delete", ep.PostPaymentDelete)
	g.POST("/payment/reorder", ep.PostPaymentReorder)
}

func (ep *Order) Get(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.orderService.Fetch(c.Request().Context(), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OrderHeaderW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Order) Post(c echo.Context) error {
	req := &model.OrderHeaderW2FormRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	req.Record.ID = req.RecID
	err := ep.orderService.SaveForm(c.Request().Context(), req.Record)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OrderHeaderW2FormResponse{Status: "success", RecID: req.RecID})
}

func (ep *Order) PostDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.orderService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrOrderDeleteProtect) {
		return c.JSON(http.StatusUnprocessableEntity, model.ResponseErrorJson("You can delete only 20 orders per request"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Order) GetCounter(c echo.Context) error {
	count, err := ep.orderService.FetchCounter(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(count))
}

func (ep *Order) GetLine(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, summary, err := ep.orderService.FetchLines(c.Request().Context(), int64(orderID), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OrderLineW2GridResponse{Status: "success", Records: rows, Total: count, Summary: []any{summary}})
}

func (ep *Order) GetStatus(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.orderStatusService.Fetch(c.Request().Context(), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.OrderStatusW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Order) GetStatusDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.orderStatusService.FetchDropdown(c.Request().Context(), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownBasicResponse{Status: "success", Records: rows})
}

func (ep *Order) PostStatusSave(c echo.Context) error {
	req := &model.OrderStatusW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err := ep.orderStatusService.SavePartials(c.Request().Context(), req.Changes)
	if errors.Is(err, service.ErrOrderStatusExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Status with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Order) PostStatusDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.orderStatusService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrOrderStatusHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Status contains related orders, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Order) PostStatusReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.orderStatusService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrOrderStatusInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Order) GetPayment(c echo.Context) error {
	req := &model.W2GridDataRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, count, err := ep.paymentMethodService.Fetch(c.Request().Context(), req.ToFindManyParams())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.PaymentMethodW2GridResponse{Status: "success", Records: rows, Total: count})
}

func (ep *Order) GetPaymentDropdown(c echo.Context) error {
	req := &model.W2DropdownRequest{}
	if err := json.Unmarshal([]byte(c.QueryParam("request")), req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	rows, err := ep.paymentMethodService.FetchDropdown(c.Request().Context(), req.Search, req.Max)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.W2DropdownBasicResponse{Status: "success", Records: rows})
}

func (ep *Order) PostPaymentSave(c echo.Context) error {
	req := &model.PaymentMethodW2SaveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson(err.Error()))
	}

	_, err := ep.paymentMethodService.SavePartials(c.Request().Context(), req.Changes)
	if errors.Is(err, service.ErrPaymentMethodExists) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Payment method with the same name already exists"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Order) PostPaymentDelete(c echo.Context) error {
	req := &model.W2GridRemoveRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	_, err := ep.paymentMethodService.DeleteListID(c.Request().Context(), req.ID)
	if errors.Is(err, service.ErrPaymentMethodHasRelated) {
		return c.JSON(http.StatusConflict, model.ResponseErrorJson("Payment method contains related orders, please remove them first"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}

func (ep *Order) PostPaymentReorder(c echo.Context) error {
	req := &model.W2GridReorderRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("bad request"))
	}

	err := ep.paymentMethodService.Reorder(c.Request().Context(), req.ID, req.MoveBefore)
	if errors.Is(err, service.ErrPaymentMethodInvalidReorder) {
		return c.JSON(http.StatusBadRequest, model.ResponseErrorJson("Invalid reorder parameters"))
	} else if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.ResponseSuccessJson(nil))
}
