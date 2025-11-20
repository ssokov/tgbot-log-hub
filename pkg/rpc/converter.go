package rpc

import (
	"log"
	logshub "logs-hub-backend/pkg/logs-hub"
	"strconv"
)

func newServices(services []logshub.Service) []ServiceResponse {
	var res []ServiceResponse
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

	log.Println("logService: ", logService)

	serviceResponse.ID = strconv.Itoa(logService[0].ServiceID)
	if logService[0].Service != nil {
		serviceResponse.Name = logService[0].Service.Name
	} else {
		log.Println("AaAAAAAAAAAAAA")
	}
	// serviceResponse.Name = (logService[0].Service.Name)

	var logs []Log
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
		// Params: serviceLog.User.Params,
		Date: serviceLog.CreatedAt,
	}
}
