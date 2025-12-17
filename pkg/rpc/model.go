package rpc

import "time"

type ServiceResponse struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type LogsService struct {
	Service ServiceResponse `json:"service"`
	Logs    []Log           `json:"logs"`
}

type Log struct {
	Type        string                 `json:"type" example:"error"`
	ErrorCode   int                    `json:"error_code" example:"500"`
	Text        string                 `json:"message" example:"invalid tg_id"`
	TgUserID    int                    `json:"tg_user_id" example:"123456789"`
	TgNickname  string                 `json:"tg_nickname,omitempty" example:"john_doe"`
	ServiceName string                 `json:"service_name,omitempty" example:"my-telegram-bot"`
	Date        time.Time              `json:"timestamp" example:"2025-10-22T12:34:56Z"`
}

type LogReq struct {
	Type        string                 `json:"type" example:"error"`
	Text        string                 `json:"message" example:"invalid tg_id"`
	TgUserID    int                    `json:"tg_user_id" example:"123456789"`
	TgNickname  string                 `json:"tg_nickname,omitempty" example:"john_doe"`
	ServiceName string                 `json:"service_name,omitempty" example:"my-telegram-bot"`
}
