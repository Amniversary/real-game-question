package logger

import (
	"context"
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro/client"
)

type ClientLogWrapper struct {
	client.Client
}

func (l *ClientLogWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Logf("[wrapper][%v] client request to service: %s method: %s\n", time.Now().Format("2006-01-02 15:04:05"), req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// implements client.Wrapper as logWrapper
func ClientLogWrap(c client.Client) client.Client {
	return &ClientLogWrapper{c}
}
