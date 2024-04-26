package request

import (
	"errors"
)

var (
	ErrParsingPathFailed   = errors.New("request: parsing path failed")
	ErrDecodingBase64Body  = errors.New("request: decoding base64 body")
	ErrFailToCreateRequest = errors.New("request: fail to create request")
)
