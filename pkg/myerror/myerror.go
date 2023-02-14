package myerror

// BadRequestError HTTP Status Code: 400
type BadRequestError struct {
	Err error
}

func (e *BadRequestError) Error() string {
	return "Bad Request Error"
}

// InternalServerError HTTP Status Code: 500
type InternalServerError struct {
	Err error
}

func (e *InternalServerError) Error() string {
	return "Internal Server Error"
}
