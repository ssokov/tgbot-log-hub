package handlers

import "github.com/labstack/echo/v4"

type ServiceHandler struct {
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}

func (h *ServiceHandler) Register(c echo.Context) error {
	return nil
}
func (h *ServiceHandler) GetApiKey(c echo.Context) error {
	return nil
}
func (h *ServiceHandler) GetServices(c echo.Context) error {
	return nil
}
func (h *ServiceHandler) UpdateService(c echo.Context) error {
	return nil
}
func (h *ServiceHandler) DeleteService(c echo.Context) error {
	return nil
}
func (h *ServiceHandler) GetLog(c echo.Context) error {
	return nil
}
func (h *ServiceHandler) GetLogByFilter(c echo.Context) error {
	return nil
}
