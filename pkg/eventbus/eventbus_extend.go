package eventbus

//-----------------------------------------
type DispatchInfos []DispatchInfo
type DispatchInfo struct {
	EventName string
	Handle    EventHandler //ctx.bind  handler,recoverry    || return data
}

type Dispatcher interface {
	Dispatcher() DispatchInfos
}
