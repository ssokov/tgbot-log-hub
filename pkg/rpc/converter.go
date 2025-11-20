package rpc

import (
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
	return Log{
		Type:      serviceLog.Type.TypeName,
		ErrorCode: *serviceLog.ErrorCode,
		Text:      *serviceLog.Message,
		TgUserID:  int(*serviceLog.UserID),
		// Params:    serviceLog.User.Params,
		Date: serviceLog.CreatedAt,
	}
}
