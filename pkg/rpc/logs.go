package rpc

import (
	"context"
	"logs-hub-backend/pkg/db"
	logshub "logs-hub-backend/pkg/logs-hub"

	"github.com/vmkteam/embedlog"
	"github.com/vmkteam/zenrpc/v2"
)

type LogService struct {
	zenrpc.Service
	embedlog.Logger

	logManager *logshub.LogManager
}

func NewLogService(dbc db.DB, logger embedlog.Logger) *LogService {
	return &LogService{
		Logger:     logger,
		logManager: logshub.NewLogManager(dbc, logger),
	}
}

func (ls *LogService) Get(ctx context.Context) ([]ServiceResponse, error) {
	services, err := ls.logManager.Get(ctx)
	if err != nil {
		return nil, err
	}

	return newServices(services), nil
}

func (ls *LogService) GetLogsByServiceID(ctx context.Context, serviceID int) (LogsService, error) {

	serviceLogs, err := ls.logManager.GetLogsService(ctx, serviceID)
	if err != nil {
		return LogsService{}, err
	}

	return newLogServices(serviceLogs), err
}

