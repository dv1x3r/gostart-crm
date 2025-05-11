package sse

import "github.com/labstack/echo/v4"

func Flush(w *echo.Response, event Event) error {
	if err := event.MarshalTo(w); err != nil {
		return err
	}
	w.Flush()
	return nil
}
