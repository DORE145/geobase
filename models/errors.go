package models

// NotFoundResponse is an error that is used when the location is not found
// swagger:response notFound
type NotFoundResponse struct {
	// The error message
	// in: body
	Message string `json:"message"`
}

// BadRequestResponse is an error that is used when required query parameter is not found
// swagger:response badParameter
type BadRequestResponse struct {
	// The error message
	// in: body
	Message string `json:"message"`
}
