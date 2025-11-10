package app

import (
	"context"
	"time"

	"logs-hub-backend/pkg/db"
	"logs-hub-backend/pkg/http"

	"github.com/go-pg/pg/v10"
	monitor "github.com/hypnoglow/go-pg-monitor"
	"github.com/labstack/echo/v4"
	"github.com/vmkteam/embedlog"
)

type Config struct {
	Database *pg.Options
	Server   struct {
		Host    string
		Port    int
		IsDevel bool
	}
	Sentry struct {
		Environment string
		DSN         string
	}
}

type App struct {
	embedlog.Logger
	appName string
	cfg     Config
	db      db.DB
	dbc     *pg.DB
	mon     *monitor.Monitor
	echo    *echo.Echo
}

func New(appName string, sl embedlog.Logger, cfg Config, db db.DB, dbc *pg.DB) *App {
	a := &App{
		appName: appName,
		cfg:     cfg,
		db:      db,
		dbc:     dbc,
		echo:    http.NewRouter(),
		Logger:  sl,
	}

	return a
}

// Run is a function that runs application.
func (a *App) Run(ctx context.Context) error {
	return a.echo.Start(":8080")
}

// Shutdown is a function that gracefully stops HTTP server.
func (a *App) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	a.mon.Close()

	return a.echo.Shutdown(ctx)
}
