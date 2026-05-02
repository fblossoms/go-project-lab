package exception

import (
	"errors"

	"github.com/infraboard/mcube/v2/tools/pretty"
)

func NewApiException(code int, message string) *ApiException {
	return &ApiException{
		Code:    code,
		Message: message,
	}
}

// ApiException 用于描述业务异常
// 实现自定义异常
type ApiException struct {
	// 自定义业务异常的编码，50001 表示Token过期
	Code int `json:"code"`
	// 异常的描述信息
	Message string `json:"message"`
	// 不会出现在Body里面，序列化成Json，HTTP response 进行set
	HttpCode int `json:"-"`
}

func (e *ApiException) Error() string {
	return e.Message
}

func (e *ApiException) String() string {
	//dj, _ := json.MarshalIndent(e, "", "  ")
	//return string(dj)
	return pretty.ToJSON(e)
}

func (e *ApiException) WithMessage(msg string) *ApiException {
	e.Message = msg
	return e
}

// IsApiException 通过Code来比较错误
func IsApiException(err error, code int) bool {
	var apiErr *ApiException
	if errors.As(err, &apiErr) {
		return apiErr.Code == code
	}
	return false
}
