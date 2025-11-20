package logs_hub

import "logs-hub-backend/pkg/db"

func newServices(services []db.Service) []Service {
	res := make([]Service, 0, len(services))
	for _, service := range services {
		res = append(res, newService(service))
	}
	return res
}

func newService(service db.Service) Service {
	return Service(service)
}

func newLogServices(serviceLogs []db.ServiceLog) []ServiceLog {
	res := make([]ServiceLog, 0, len(serviceLogs))
	for _, servserviceLog := range serviceLogs {
		res = append(res, newLogService(servserviceLog))
	}
	return res
}

func newLogService(serviceLog db.ServiceLog) ServiceLog {
	return ServiceLog(serviceLog)
}
