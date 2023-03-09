package eventbus

//-----------------------------------------
type DispatchInfos []DispatchInfo
type DispatchInfo struct {
	EventName string
	Handle    any 
}

type Dispatcher interface {
	Dispatcher() DispatchInfos
}
