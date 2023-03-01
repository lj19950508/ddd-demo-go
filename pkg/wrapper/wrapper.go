package wrapper


import (
	"github.com/lj19950508/ddd-demo-go/pkg/errors"
	"github.com/lj19950508/ddd-demo-go/internal/adapter/in/web/dto"
)

//bizcode 0 正常
//bizcode -1 异常
//bizCode ++ 业务码

func Error(err error) (int, *dto.Result[any]) {
	//经过一条[err,httpStatus,bizCode] handler 处理  返回
	// err ->httpStatus,bizCode
	// ErrInvalidTransaction = errors.New("invalid transaction")
	//预料之外的都是500
	//for erroris  return  ,, default reutrn -1,-1  (没有注册该异常)
	bizCode := pkg.SearchErr(err)
	if bizCode == -1 {
		//野生异常 直接被panic处理了
		//自己创造出来的异常，或者第三方库生成的异常可以搜集到，如果是第三方库的异常 要在生成的时候去根据情况处理,如果是自己创造出来的异常则不用管直接panic
		//第三方库的异常如何处理？
		//1.意料之中的直接处理
		//2.意料之外的直接panic?
		panic(err)
		//这里panic 是没有找到对应的异常的情况 就是 程序中如果有野生的异常还被return到这的 就只能panic ， error生成控制好就不会存在

	}

	return 400, dto.NewResult[any](nil, bizCode, err.Error())
}

func ResultData(data any) (int,*dto.Result[any]) {

	return 200, dto.NewResult(data, 0, "")
}

func ResultMsg(msg string) (int, *dto.Result[any]) {

	return 200, dto.NewResult[any](nil, 0, msg)
}

func Result(msg string, data any) (int, *dto.Result[any]) {

	return 200, dto.NewResult(data, 0, "")
}
