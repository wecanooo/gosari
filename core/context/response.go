package context

import (
	"net/http"

	"github.com/wecanooo/gosari/core/constants"
	"github.com/wecanooo/gosari/core/errno"
	"github.com/wecanooo/gosari/core/pkg/serializer"
)

type CommonResponse struct {
	Code    constants.LogicCode `json:"code"`
	Message string              `json:"msg"`
	Data    interface{}         `json:"data,omitempty"`
}

func NewCommonResponse(code constants.LogicCode, message string, data interface{}) *CommonResponse {
	return &CommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse(message string, data interface{}) *CommonResponse {
	return NewCommonResponse(constants.SuccessCode, message, data)
}

func NewErrResponse(e *errno.Errno) *CommonResponse {
	return NewCommonResponse(e.Code, e.Message, e.Data)
}

func (c *AppContext) SuccessJSON(data interface{}) error {
	return c.JSON(http.StatusOK, NewSuccessResponse("ok", data))
}

func (c *AppContext) ErrorJSON(e *errno.Errno) error {
	return c.JSON(e.HTTPCode, NewErrResponse(e))
}

func (c *AppContext) SerializeData(data interface{}) error {
	d := serializer.Serialize(data)
	return c.SuccessJSON(d)
}
