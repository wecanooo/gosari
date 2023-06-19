package constants

type (
	LogicCode int
)

const (
	SuccessCode       LogicCode = 0
	UnknownErrorCode  LogicCode = -1
	RequestErrorCode  LogicCode = 100
	ResourceErrorCode LogicCode = 101
	DatabaseErrorCode LogicCode = 102
	TokenErrorCode    LogicCode = 103
	NotFoundErrorCode LogicCode = 104
	AuthErrorCode     LogicCode = 105
)
