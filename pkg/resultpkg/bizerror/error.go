package bizerror

import (
	"errors"
)

//只存在业务错误  err->bizCode
// i18n错误

type BizError struct{
	error
	BizCode int
}

func NewBizError(bizCode int,msg string) error{
	return &BizError{
		errors.New(msg),
		bizCode,
	}
}



