package utils

import (
	"os"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

var (
	logOnce sync.Once
	log     zerolog.Logger
)

func GetLogger() zerolog.Logger {
	logOnce.Do(func() {
		writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		log = zerolog.New(writer).With().Timestamp().Logger()
	})

	return log
}

func LogEchoRequestFunc(c echo.Context, v middleware.RequestLoggerValues) error {
	var event *zerolog.Event
	if v.Error == nil {
		event = log.Info()
	} else if v.Status >= 500 {
		event = log.Error()
	} else {
		event = log.Warn()
	}
	event.
		Str("id", v.RequestID).
		Int("status", v.Status).
		Str("remote_ip", v.RemoteIP).
		Dur("latency", v.Latency).
		Str("host", v.Host).
		Str("method", v.Method).
		Str("uri", v.URI).
		Err(v.Error).
		Msg("echo")
	return nil
}

func LogEchoRecoverFunc(c echo.Context, err error, stack []byte) error {
	requestID := c.Response().Header().Get(echo.HeaderXRequestID)
	log.Error().
		Str("id", requestID).
		Err(err).
		Bytes("stack", stack).
		Msg("recover")
	return err
}
