package discount

import (
	"context"
	"encoding/json"
	"fmt"
	"gostart-crm/internal/lib/webapi/sse"
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Register(g *echo.Group) {
	g.GET("/random/discounts", func(c echo.Context) error {
		ctx := c.Request().Context()
		sse.SetHeaderEntries(c.Response().Header())
		discounts := getDiscountsChan(ctx)

		for {
			select {
			case <-ctx.Done():
				return nil

			case res := <-discounts:
				if res.Error != nil {
					return res.Error
				}

				if err := sse.Flush(c.Response(), sse.Event{Data: res.Data}); err != nil {
					return err
				}

				c.Logger().Info("/random/discounts sse data sent, ip: ", c.RealIP())
			}
		}
	})
}

type Discount struct {
	Value string `json:"value"`
}

type DiscountsChanResult struct {
	Data  []byte
	Error error
}

func getDiscountsChan(ctx context.Context) <-chan DiscountsChanResult {
	ch := make(chan DiscountsChanResult)
	go func() {
		defer close(ch)

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				discounts := map[int]Discount{}
				for i := range 250 {
					discounts[i] = Discount{Value: fmt.Sprintf("%.0f%%", rand.Float64()*100)}
				}

				data, err := json.Marshal(discounts)
				ch <- DiscountsChanResult{Data: data, Error: err}
			}
		}
	}()
	return ch
}
