package user

var (
	// ErrNotFound means the item was not found
	ErrNotFound = "not found"
	// ErrNotFoundDesc describes the not found state
	ErrNotFoundDesc = "Provided email does not exist"
	// ErrRequestBodyCode is the code for error in request body
	ErrRequestBodyCode = "request_body_error"
	// ErrRequestBodyDesc is the description for error in request body
	ErrRequestBodyDesc = "Invalid request body"
)
