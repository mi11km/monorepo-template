// Package api contains the entrypoint for the API server.
package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"connectrpc.com/connect"
	v1 "github.com/mi11km/monorepo-template/go/apps/sample/infrastructure/rpc/v1"
	"github.com/mi11km/monorepo-template/go/apps/sample/infrastructure/rpc/v1/v1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Run initializes and starts the API server.
func Run(ctx context.Context) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()
	path, handler := v1connect.NewHealthServiceHandler(&healthServiceHandler{})
	mux.Handle(path, handler)
	svr := &http.Server{
		Addr:              ":8080",
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: time.Second,
	}

	go func() {
		defer stop()

		err := svr.ListenAndServe()
		if err != nil {
			slog.Error("server: failed to listen and serve", slog.Any("err", err))
		}
	}()

	<-ctx.Done()

	err := svr.Shutdown(ctx)
	if err != nil {
		slog.Error("server: failed to shutdown", slog.Any("err", err))
	}
}

var _ v1connect.HealthServiceHandler = (*healthServiceHandler)(nil)

type healthServiceHandler struct{}

func (h *healthServiceHandler) Check(
	_ context.Context,
	_ *connect.Request[v1.CheckRequest],
) (*connect.Response[v1.CheckResponse], error) {
	return connect.NewResponse(&v1.CheckResponse{Status: v1.ServingStatus_SERVING_STATUS_OK}), nil
}

func (h *healthServiceHandler) Watch(
	_ context.Context,
	_ *connect.Request[v1.WatchRequest],
	stream *connect.ServerStream[v1.WatchResponse],
) error {
	err := stream.Send(&v1.WatchResponse{Status: v1.ServingStatus_SERVING_STATUS_OK})
	if err != nil {
		return fmt.Errorf("watch: %w", err)
	}

	return nil
}
