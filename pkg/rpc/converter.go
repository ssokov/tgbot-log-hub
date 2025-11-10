package rpc

import (
	logshub "logs-hub-backend/pkg/logs-hub"
	"strconv"
)

func newServices(services []logshub.Service) []Service {
	var res []Service
	for _, service := range services {
		res = append(res, newService(service))
	}
	return res
}

func newService(service logshub.Service) Service {
	return Service{
		ID:   strconv.Itoa(service.ID),
		Name: service.Name,
	}
}
