package models

import (
	"github.com/labstack/echo/v4"
	"github.com/phanorcoll/go_mongo/pkg/domain"
)

// RegisterRequest struct    represents the type the API accepts
// for registering a new user.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequest struct    represents the types the API accepts
// for logging a user.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ValidateRegisterRequest function    will run a number of validations
// when a user tries to register
func ValidateRegisterRequest(c echo.Context) (*domain.User, *Error) {
	registerRequest := new(RegisterRequest)
	if err := c.Bind(registerRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(registerRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")
	}

	if len(registerRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")
	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: registerRequest.Username,
		Password: registerRequest.Password,
	}, nil
}

// ValidateLoginRequest function    will run a number of validations
// when a user tries to login
func ValidateLogInRequest(c echo.Context) (*domain.User, *Error) {
	loginRequest := new(LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return nil, BindError()
	}

	var validationErrors []string

	if len(loginRequest.Password) < 8 {
		validationErrors = append(validationErrors, "Password must be 8 characters")
	}

	if len(loginRequest.Username) < 3 {
		validationErrors = append(validationErrors, "Username must be longer than 2 characters")
	}

	if len(validationErrors) > 0 {
		return nil, ValidationError(validationErrors)
	}

	return &domain.User{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}, nil
}
