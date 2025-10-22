package handlers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct{}

type jwtCustomClaims struct {
	Login string `json:"login"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type RegisterRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
	Email    string `json:"email" example:"misha.sokovih@gmail.com"`
}

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}

// Register godoc
// @Summary Register a new admin
// @Description Create a new admin account
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Admin registration data"
// @Success 201 {object} map[string]string "Successfully registered"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /admin/register [post]
func (h *AdminHandler) Register(c echo.Context) error {
	return nil
}

// Login godoc
// @Summary Admin login
// @Description Login with username and password to receive JWT token
// @Tags Admin
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Admin login data"
// @Success 200 {object} LoginResponse "JWT token returned"
// @Failure 400 {object} ErrorResponse "Invalid credentials"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /admin/login [post]
func (h *AdminHandler) Login(c echo.Context) error {
	return nil
}

// Logout godoc
// @Summary Admin logout
// @Description Invalidate admin JWT token (if token tracking is implemented)
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Successfully logged out"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /admin/logout [post]
func (h *AdminHandler) Logout(c echo.Context) error {
	return nil
}
