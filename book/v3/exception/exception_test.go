package exception_test

import (
	"go18/book/v3/exception"
	"testing"
)

func CheckIsError() error {
	return exception.ErrNotFound("book %d not found", 1)
}

func TestException(t *testing.T) {
	err := CheckIsError()
	t.Log(err)

	// 断言接口类型
	v, ok := err.(*exception.ApiException)
	if ok {
		t.Log(v.Code)
		t.Log(v.String())
	}

	t.Log(exception.IsApiException(err, exception.CODE_NOT_FOUND))
}
