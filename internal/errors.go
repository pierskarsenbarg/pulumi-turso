package turso

import (
	"errors"
	"fmt"
)

type errorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err *errorResponse) Error() string {
	return fmt.Sprintf("%s API error: %s", err.Code, err.Message)
}

func GetErrorStatusCode(err error) string {
	var errResp *errorResponse
	if errors.As(err, &errResp) {
		return errResp.Code
	}
	return "0"
}
