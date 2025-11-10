package logs_hub

import "logs-hub-backend/pkg/db"

func newServices(services []db.Service) []Service {
	var res []Service
	for _, service := range services {
		res = append(res, newService(service))
	}
	return res
}

func newService(service db.Service) Service {
	return Service(service)
}
