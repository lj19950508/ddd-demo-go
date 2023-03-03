package resultpkg

import (
	"fmt"
	"net/http"
	"github.com/lj19950508/ddd-demo-go/pkg/resultpkg/bizerror"
)

type Result struct {
	Data    any    `json:"data"`
	BizCode int    `json:"bizCode"`
	Msg     string `json:"msg"`
}

func NewResult(data any, bizCode int, msg string) *Result {
	return &Result{
		Data:    data,
		BizCode: bizCode,
		Msg:     msg,
	}
}

var (
	//正常业务码
	BizCodeNormal = 0
	//找不到当前异常
	BizCodeError = -1
)

func Ok(data any) *Result {
	return NewResult(data, BizCodeNormal, "")
}

func OkMsg(msg string) *Result {
	return NewResult(nil, BizCodeNormal, msg)
}

func Fail(msg string)* Result{
	return NewResult(nil, BizCodeError, msg)
}


func Error(err error) (int, *Result) {
	//error is biz error
	switch v := err.(type) {
	case *bizerror.BizError:
		return http.StatusOK, NewResult(nil, v.BizCode, err.Error())
	default:
		return http.StatusInternalServerError, NewResult(nil, BizCodeError, fmt.Sprintf("%+v", err))
	}


	
	//业务异常不包装堆栈
}