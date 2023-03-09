package eventbus

type Event struct {
	Payload      any    `json:"payload"`
	Id           int64  `json:"id"`
	Name         string `json:"event"`
	ExcuteResult any    `json:"result"`
}

func NewEvent(Id int64, Name string, Payload any) *Event {
	return &Event{
		Id:      Id,
		Payload: Payload,
		Name:    Name,
	}
}

type EventBus interface {
	Publish(evt *Event) error //receive result  bindreuslt
	Subscribe(evt *Event) error
}

type DispatchInfos []DispatchInfo
type DispatchInfo struct {
	EventName string
	Handle    func(evt *Event) //ctx.bind  handler,recoverry    || return data
}

type Dispatcher interface {
	Dispatcher() DispatchInfos
}
