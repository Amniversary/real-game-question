package logger

import (
	"context"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/server"
)

func ServerLogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[%v] server request: %s", time.Now().Format("2006-01-02 15:04:05"), req.Method())
		return fn(ctx, req, rsp)
	}
}
