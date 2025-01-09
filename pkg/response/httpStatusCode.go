package response

const (
	ErrorCodeSuccess      = 2001 // Success
	ErrorCodeParamInvalid = 2003 // Email is invalid
	ErrorInvalidToken     = 3001 // invalid token
)

// msg
var msg = map[int]string{
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email is invalid",
	ErrorInvalidToken:     "invalid token",
}
