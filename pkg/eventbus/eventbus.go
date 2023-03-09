package eventbus

//TODO 这个event承载太多了 或许可以试试用options方式gg
type Event struct {
	Payload      any    `json:"payload"`
	Response     any    `json:"response"`
	Name         string `json:"event"`
}

func NewEvent(Name string, Payload any) *Event {
	return &Event{
		Payload: Payload,
		Name:    Name,
	}
}


//----------------------------------------------
type EventBus interface {
	Publish(evt *Event) error //receive result  bindreuslt	
	Subscribe(evt string,handler any) error
}


