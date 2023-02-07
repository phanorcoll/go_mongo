// Package models provides models  
// Will contain internal model structs used between application
// layers and general purpose usage.
package models

// Error struct    represents and store information for any errors
// we have as part of the business rules
type Error struct {
	Message    string   `json:"message"`
	Code       int      `json:"code"`
	Name       string   `json:"name"`
	Error      error    `json:"-"`
	Validation []string `json:"validation,omitempty"`
}

// BindError function    will throw a common error when trying to bind
// the body of a request to unmarshall the body.
func BindError() *Error {
	return &Error{
		Code:    400,
		Message: "Error processing request",
		Name:    "BiND_ERROR",
	}
}

// ValidationError function    will throw an error depending on the request.
func ValidationError(errors []string) *Error {
	return &Error{
		Code:       400,
		Name:       "VALIDATION",
		Message:    "A validation error occurred",
		Validation: errors,
	}
}
