package eventbus

type Event struct {
	Payload      any    `json:"payload"`
	Id           int64  `json:"id"`
	Name         string `json:"event"`
	Return 	     bool   `json:"return"`
	ExcuteResult any    `json:"result"`
	Compensation string `json:"compensation"` //补偿队列的名字
}

func NewEvent(Id int64, Name string, Payload any) *Event {
	return &Event{
		Id:      Id,
		Payload: Payload,
		Name:    Name,
	}
}
//----------------------------------------------

type EventBus interface {
	Publish(evt *Event) error //receive result  bindreuslt
	Subscribe(dispatcher Dispatcher) error
}

type DispatchInfos []DispatchInfo
type DispatchInfo struct {
	EventName string
	Handle    EventHandler //ctx.bind  handler,recoverry    || return data
}

type Dispatcher interface {
	Dispatcher() DispatchInfos
}
type EventHandler func(evt *Event) 
