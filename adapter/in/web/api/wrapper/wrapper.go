package wrapper

import (
	"fmt"
	"net/http"

	"github.com/lj19950508/ddd-demo-go/domain/bizerror"
	dto "github.com/lj19950508/ddd-demo-go/pkg/common_dto"
)

//bizcode 0 正常
//bizcode -1 异常
//bizCode ++ 业务码

// var (
// 	BizCodeNormal = errors.New("order status error")
// )

func Error(err error) (int, *dto.Result) {
	bizCode := bizerror.SearchErr(err)
	if bizCode == bizerror.BizCodeError {
		//这里是野生的异常，比如数据库错误等这些 ,这些会包装他的堆栈以便于追踪  (调用第三方库的第一层的错误要包装 如 repository的dberror)
		return http.StatusInternalServerError, dto.NewResult(nil, bizCode, fmt.Sprintf("%+v", err))
	}
	//业务异常不包装堆栈
	return http.StatusOK, dto.NewResult(nil, bizCode, err.Error())
}

//

func ResultData(data any) *dto.Result {

	return dto.NewResult(data, bizerror.BizCodeNormal, "")
}

func ResultMsg(msg string) *dto.Result {

	return dto.NewResult(nil, bizerror.BizCodeNormal, msg)
}

func Result(msg string, data any) *dto.Result {

	return dto.NewResult(data, bizerror.BizCodeNormal, "")
}
