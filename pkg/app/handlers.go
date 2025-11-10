package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/vmkteam/appkit"
	"github.com/vmkteam/rpcgen/v2"
	"github.com/vmkteam/rpcgen/v2/golang"
	zm "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"logs-hub-backend/pkg/rpc"
	"net/http"
)

// runHTTPServer is a function that starts http listener using labstack/echo.
func (a *App) runHTTPServer(ctx context.Context, host string, port int) error {
	listenAddress := fmt.Sprintf("%s:%d", host, port)
	addr := "http://" + listenAddress
	a.Print(ctx, "starting http listener", "url", addr)

	a.registerAPIHandlers(listenAddress)

	return a.echo.Start(listenAddress)
}

// registerAPIHandlers registers main rpc server.
func (a *App) registerAPIHandlers(listenAddress string) {
	srv := rpc.New(a.db, a.Logger, a.cfg.Server.IsDevel)
	gen := rpcgen.FromSMD(srv.SMD())

	a.echo.Any("/v1/rpc/", zm.EchoHandler(zm.XRequestID(srv)))
	a.echo.Any("/v1/rpc/doc/", appkit.EchoHandlerFunc(zenrpc.SMDBoxHandler))
	a.echo.Any("/v1/rpc/openrpc.json", appkit.EchoHandlerFunc(rpcgen.Handler(gen.OpenRPC("logshub", listenAddress+"/v1/rpc"))))
	a.echo.Any("/v1/rpc/api.go", appkit.EchoHandlerFunc(rpcgen.Handler(gen.GoClient(golang.Settings{Package: "logshub"}))))
}

// registerDebugHandlers adds /debug/pprof handlers into a.echo instance.
func (a *App) registerDebugHandlers() {
	dbg := a.echo.Group("/debug")

	// add pprof integration
	dbg.Any("/pprof/*", appkit.PprofHandler)

	// add healthcheck
	a.echo.GET("/status", func(c echo.Context) error {
		// test postgresql connection
		err := a.db.Ping(c.Request().Context())
		if err != nil {
			a.Error(c.Request().Context(), "failed to check db connection", "err", err)
			return c.String(http.StatusInternalServerError, "DB error")
		}
		return c.String(http.StatusOK, "OK")
	})

	// show all routes in devel mode
	if a.cfg.Server.IsDevel {
		a.echo.GET("/", appkit.RenderRoutes(a.appName, a.echo))
	}
}
