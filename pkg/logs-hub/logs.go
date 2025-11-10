package logs_hub

import (
	"context"
	"github.com/vmkteam/embedlog"
	"logs-hub-backend/pkg/db"
)

type LogManager struct {
	dbc    db.DB
	tlRepo db.TgbotLogHubRepo
	embedlog.Logger
}

func (m LogManager) Get(ctx context.Context) ([]Service, error) {
	services, err := m.tlRepo.ServicesByFilters(ctx, &db.ServiceSearch{}, db.PagerNoLimit)
	if err != nil {
		return nil, err
	}

	return newServices(services), err

}

func NewLogManager(dbc db.DB, logger embedlog.Logger) *LogManager {
	return &LogManager{dbc: dbc, Logger: logger, tlRepo: db.NewTgbotLogHubRepo(dbc)}
}
