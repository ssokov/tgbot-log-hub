package rpc

import (
	"logs-hub-backend/pkg/db"
	logshub "logs-hub-backend/pkg/logs-hub"
	"strconv"
)

func newServices(services []logshub.Service) []ServiceResponse {
	res := make([]ServiceResponse, 0, len(services))
	for _, service := range services {
		res = append(res, newService(service))
	}
	return res
}

func newService(service logshub.Service) ServiceResponse {
	return ServiceResponse{
		ID:   strconv.Itoa(service.ID),
		Name: service.Name,
	}
}

func newLogServices(logService []logshub.ServiceLog) LogsService {
	var serviceResponse ServiceResponse

	if len(logService) == 0 {
		return LogsService{}
	}

	serviceResponse.ID = strconv.Itoa(logService[0].ServiceID)
	serviceResponse.Name = logService[0].Service.Name
	logs := make([]Log, 0, len(logService))

	for _, serviveLog := range logService {
		logs = append(logs, newLog(serviveLog))
	}

	return LogsService{
		Service: serviceResponse,
		Logs:    logs,
	}
}

func newLog(serviceLog logshub.ServiceLog) Log {

	log := Log{
		Type: serviceLog.Type.TypeName,
		Date: serviceLog.CreatedAt,
	}

	if serviceLog.ErrorCode != nil {
		log.ErrorCode = *serviceLog.ErrorCode
	} else {
		log.ErrorCode = -1
	}

	if serviceLog.Message != nil {
		log.Text = *serviceLog.Message
	}

	if serviceLog.UserID != nil {
		log.TgUserID = serviceLog.User.TgID
	} else {
		log.TgUserID = -1
	}
	return log
}

func newServiceLog(log LogReq) logshub.ServiceLog {
	// userID := int64(log.TgUserID)
	DemoID := int64(25)

	return logshub.ServiceLog{
		Message: &log.Text,
		UserID:  &DemoID,
		User: &db.ServiceUser{
			ID:       int(DemoID),
			TgID:     log.TgUserID,
			Nickname: &log.TgNickname,
		},
	}
}
