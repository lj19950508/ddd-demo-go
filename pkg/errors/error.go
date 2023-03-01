package pkg

import "errors"


//只存在业务错误  err->bizCode
// i18n错误


var (
	BizCodeNormal = 0
	BizCodeError  =-1
)

var (
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
