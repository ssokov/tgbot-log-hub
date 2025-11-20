package logs_hub

import (
	"context"
	"logs-hub-backend/pkg/db"

	"github.com/vmkteam/embedlog"
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

func (m LogManager) GetLogsService(ctx context.Context, serviceID int) ([]ServiceLog, error) {
	logs, err := m.tlRepo.ServiceLogsByFilters(ctx, &db.ServiceLogSearch{ServiceID: &serviceID}, db.PagerNoLimit, m.tlRepo.FullServiceLog())

	if err != nil {
		m.Logger.Errorf("GetLogsService: failed to get logs by service id=%d: %v", serviceID, err)
		return nil, err
	}

	return newLogServices(logs), err
}
