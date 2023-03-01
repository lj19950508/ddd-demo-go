package wrapper

import "errors"

// import "net/http"

//i18n 

// msg

// err

//只存在业务错误  err->bizCode
var (
	ErrVar = errors.New("testBizError")
)

var bizErrorMap = map[error]int{
	ErrVar: 1,
}

func searchErr(err error) int {
	for target, bizCode := range bizErrorMap {
		if errors.Is(err, target) {
			return bizCode
		}
	}
	return -1
	//找不到异常直接报错
}
