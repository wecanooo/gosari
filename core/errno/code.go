package errno

import (
	"net/http"

	"github.com/wecanooo/gosari/core/constants"
)

var (
	UnknownErr  = &Errno{HTTPCode: http.StatusInternalServerError, Code: constants.UnknownErrorCode, Message: "unknown error"}
	ReqErr      = &Errno{HTTPCode: http.StatusBadRequest, Code: constants.RequestErrorCode, Message: "request error"}
	ResourceErr = &Errno{HTTPCode: http.StatusNotFound, Code: constants.ResourceErrorCode, Message: "resource error"}
	DatabaseErr = &Errno{HTTPCode: http.StatusInternalServerError, Code: constants.DatabaseErrorCode, Message: "database error"}
	TokenErr    = &Errno{HTTPCode: http.StatusNotAcceptable, Code: constants.TokenErrorCode, Message: "token error"}
	NotFoundErr = &Errno{HTTPCode: http.StatusNotFound, Code: constants.NotFoundErrorCode, Message: "route not found"}
	AuthErr     = &Errno{HTTPCode: http.StatusUnauthorized, Code: constants.AuthErrorCode, Message: "auth error"}
)
