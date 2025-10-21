package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
}

type jwtCustomClaims struct {
	Login string `json:"login"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

func (h *AdminHandler) Register(c echo.Context) error {

	return nil
}

func (h *AdminHandler) Login(c echo.Context) error {

	return nil
}

func (h *AdminHandler) Logout(c echo.Context) error {

	return nil
}
