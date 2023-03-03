package dto

type Result struct{
	Data any `json:"data"`
	BizCode int `json:"bizCode"`
	Msg string `json:"msg"`
}

func NewResult(data any,bizCode int,msg string) *Result{
	return &Result{
		Data:data,
		BizCode: bizCode,
		Msg:msg,
	}
}

