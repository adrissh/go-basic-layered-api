package appErrors

import (
	"errors"
)

// Error Sentinels untuk aplikasi Anda
var (
	ErrEmployeeNotFound      = errors.New("employee with the specified ID was not found")
	ErrEmployeeIsRequired    = errors.New("Employee ID is a required field")
	ErrEmployeeAlreadyExists = errors.New("employee with the same ID already exists")
	ErrInvalidIdEmployee     = errors.New("invalid id employe format")
	ErrUnauthorizedAccess    = errors.New("unauthorized access")
)
