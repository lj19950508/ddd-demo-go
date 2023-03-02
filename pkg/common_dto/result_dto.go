package dto

type Result[T any] struct{
	Data T `json:"data"`
	BizCode int `json:"bizCode"`
	Msg string `json:"msg"`
}

func NewResult[T any](data T,bizCode int,msg string) *Result[T]{
	return &Result[T]{
		Data:data,
		BizCode: bizCode,
		Msg:msg,
	}
}

