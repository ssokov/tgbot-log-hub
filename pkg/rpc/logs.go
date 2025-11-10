package rpc

import (
	"context"
	"github.com/vmkteam/embedlog"
	"github.com/vmkteam/zenrpc/v2"
	"logs-hub-backend/pkg/db"
	logshub "logs-hub-backend/pkg/logs-hub"
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

func (ls *LogService) Get(ctx context.Context) ([]Service, error) {
	services, err := ls.logManager.Get(ctx)
	if err != nil {
		return nil, err
	}

	return newServices(services), nil
}
