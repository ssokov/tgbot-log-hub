package app

import (
	"context"
	"fmt"
)

// runHTTPServer is a function that starts http listener using labstack/echo.
func (a *App) runHTTPServer(ctx context.Context, host string, port int) error {
	listenAddress := fmt.Sprintf("%s:%d", host, port)
	addr := "http://" + listenAddress
	a.Print(ctx, "starting http listener", "url", addr, "smdbox", addr+"/v1/rpc/doc/")

	return a.echo.Start(listenAddress)
}

func (a *App) registerHandlers() {
	a.echo.GET("/main", a.wm.MainHandler)
}
