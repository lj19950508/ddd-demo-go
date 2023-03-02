package bizerror

import "errors"


//只存在业务错误  err->bizCode
// i18n错误


var (
	//正常业务码
	BizCodeNormal = 0
	//找不到当前异常
	BizCodeError  =-1
)

var (
	//业务异常码
	ErrOrderStatusError = errors.New("order status error")
)

var bizErrorMap = map[error]int{
	ErrOrderStatusError: 1,
}

func SearchErr(err error) int {
	for target, bizCode := range bizErrorMap {		
		if errors.Is(err, target) {
			return bizCode
		}
	}
	return -1
	//找不到异常直接报错
}
