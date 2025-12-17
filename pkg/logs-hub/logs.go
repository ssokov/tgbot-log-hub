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

func (m LogManager) AddLog(ctx context.Context, typeName string, log ServiceLog) error {
	m.Logger.Printf("AddLog: adding log of type: %v", typeName)

	logType, err := m.tlRepo.OneLogType(ctx, &db.LogTypeSearch{TypeName: &typeName})
	if err != nil {
		m.Logger.Errorf("AddLog: failed to find log type %s: %v", typeName, err)
		return err
	}

	if logType == nil {
		newLogType := &db.LogType{
			TypeName: typeName,
		}
		logType, err = m.tlRepo.AddLogType(ctx, newLogType)
		if err != nil {
			m.Logger.Errorf("AddLog: failed to create log type %s: %v", typeName, err)
			return err
		}
		m.Logger.Printf("AddLog: created new log type %s with ID=%d", typeName, logType.ID)
	}

	var userIDIntPtr *int
	if log.UserID != nil {
		userIDInt := int(*log.UserID)
		userIDIntPtr = &userIDInt
	}

	m.Logger.Printf("user id: %d", *userIDIntPtr)
	serviveUser, err := m.tlRepo.OneServiceUser(ctx, &db.ServiceUserSearch{ID: userIDIntPtr})
	if err != nil {
		m.Logger.Errorf("AddLog: failed to find user %v, %v ", log.User.TgID, err)
		return err
	}

	if serviveUser == nil {
		newUser := db.ServiceUser{
			ID:       *userIDIntPtr,
			TgID:     *userIDIntPtr,
			Nickname: log.User.Nickname,
		}
		serviveUser, err = m.tlRepo.AddServiceUser(ctx, &newUser)
		if err != nil {
			m.Logger.Errorf("AddLog: failed to create user '%d': %v", userIDIntPtr, err)
			return err
		}
		m.Logger.Printf("AddLog: created new user id: %d ", serviveUser.ID)

	}

	var userID = int64(serviveUser.ID)

	dbLog := newDBServiceLog(log, logType.ID, &userID)
	_, err = m.tlRepo.AddServiceLog(ctx, &dbLog)

	if err != nil {
		m.Logger.Errorf("AddLog: failed to add service log: %v", err)
		return err
	}

	return nil
}
