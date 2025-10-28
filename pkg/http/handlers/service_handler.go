package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Service struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RegisterServiceRequest struct {
	Name string `json:"name" example:"ludomania_bot"`
	Type string `json:"type" example:"telegram_bot"`
}

type APIKeyResponse struct {
	APIKey string `json:"api_key" example:"abcd1234"`
}

type ServiceListResponse struct {
	Services []Service `json:"services"`
}

type StatusResponse struct {
	Status string `json:"status" example:"ok"`
}

type Log struct {
	Type      string                 `json:"type" example:"error"`
	ErrorCode int                    `json:"error_code" example:"500"`
	Text      string                 `json:"message" example:"invalid tg_id"`
	TgUserID  int                    `json:"tg_user_id" example:"123456789"`
	Params    map[string]interface{} `json:"params,omitempty" swaggertype:"object"`
	Date      time.Time              `json:"timestamp" example:"2025-10-22T12:34:56Z"`
}

type LogSearch struct {
	Type      string    `json:"type,omitempty" example:"error"`
	ErrorCode int       `json:"error_code,omitempty" example:"500"`
	StartDate time.Time `json:"start_date,omitempty" example:"2025-10-01T00:00:00Z"`
	EndDate   time.Time `json:"end_date,omitempty" example:"2025-10-22T23:59:59Z"`
	Text      string    `json:"text,omitempty" example:"something"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"bad request"`
}

type ServiceHandler struct {
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}

// Register godoc
// @Summary Register a new service
// @Tags Service
// @Accept json
// @Produce json
// @Param payload body RegisterServiceRequest true "Service registration payload"
// @Success 201 {object} Service
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /services/register [post]
func (h *ServiceHandler) Register(c echo.Context) error {
	return nil
}

// GetAPIKey godoc
// @Summary Get API Key for a service
// @Tags Service
// @Accept json
// @Produce json
// @Param name query string true "Service name (query param)"
// @Success 200 {object} APIKeyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /services/apikey [get]
func (h *ServiceHandler) GetAPIKey(c echo.Context) error {
	return nil
}

// GetServices godoc
// @Summary List available services
// @Tags Service
// @Accept json
// @Produce json
// @Param admin query bool false "Admin flag to include inactive/hidden services"
// @Success 200 {object} ServiceListResponse
// @Failure 500 {object} ErrorResponse
// @Router /services [get]
func (h *ServiceHandler) GetServices(c echo.Context) error {
	return nil
}

// TODO I dont know for what that method

// func (h *ServiceHandler) UpdateService(c echo.Context) error {
// 	return nil
// }

// DeleteService godoc
// @Summary Delete a service
// @Tags Service
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} StatusResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /services/{id} [delete]
func (h *ServiceHandler) DeleteService(c echo.Context) error {
	return nil
}

// GetLog godoc
// @Summary Get all logs for a service
// @Tags Service
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {array} Log
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /services/{id}/logs [get]
func (h *ServiceHandler) GetLog(c echo.Context) error {
	return nil
}

// GetLogByFilter godoc
// @Summary Get logs with filters
// @Description Get logs for a service by ID. Фильтры передаются в теле запроса.
// @Tags Service
// @Accept json
// @Produce json
// @Param id path string true "Service ID"
// @Param search body LogSearch true "Log search filters"
// @Success 200 {array} Log
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /{id}/services/logs/filter [post]
func (h *ServiceHandler) GetLogByFilter(c echo.Context) error {
	return nil
}
