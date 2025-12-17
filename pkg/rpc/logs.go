package rpc

import (
	"context"
	"logs-hub-backend/pkg/db"
	logshub "logs-hub-backend/pkg/logs-hub"

	"github.com/vmkteam/embedlog"
	"github.com/vmkteam/zenrpc/v2"
)

const (
	TelegramUpdateType = "Telegram Update"
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

func (ls LogService) Get(ctx context.Context) ([]ServiceResponse, error) {
	services, err := ls.logManager.Get(ctx)
	if err != nil {
		return nil, err
	}

	return newServices(services), nil
}

func (ls LogService) GetLogsByServiceID(ctx context.Context, serviceID int) (LogsService, error) {

	serviceLogs, err := ls.logManager.GetLogsService(ctx, serviceID)
	if err != nil {
		return LogsService{}, err
	}

	return newLogServices(serviceLogs), err
}

func (ls LogService) AddTelegramLog(ctx context.Context, log LogReq) error {

	serviceLog := newServiceLog(log)
	if err := ls.logManager.AddLog(ctx, log.Type, serviceLog); err != nil {
		ls.Logger.Errorf("AddTelegramLog: failed to add log: %v", err)
		return err
	}

	return nil
}
