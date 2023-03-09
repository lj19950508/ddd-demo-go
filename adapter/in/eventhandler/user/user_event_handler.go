package eventhandler

import (
	"errors"

	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/streadway/amqp"
)

//这个改成类似api的处理
//然后有一个eventbushandler 是
//用户子域的事件处理

//EventHandler inteface has a router  routerInfo [message,func]
//TODO eventbus need eventhandler[]  and eventbus foreach .AddEventListenr()
//   Queue=>func (evt *user.EvtUserCreate) error
// so just need eventbu

type UserEventHandler struct {
	logger   logger.Interface
}

func NewUserEventHandler(logger logger.Interface) *UserEventHandler {
	return &UserEventHandler{
		logger:logger,
	}
}

func (t *UserEventHandler) Dispatcher() eventbus.DispatchInfos {
	return eventbus.DispatchInfos{
		//默认使用user吧
		{EventName:"UserCreateEvent", Handle: t.Handler1},
		
	}
}

type Response struct{
	A int `json:"a"`
}


func (s *UserEventHandler) Handler1(d *amqp.Delivery)(any,error){
	// evt.Payload
	//如果有业务error  1.调用服务补偿  2.输出日志 
	s.logger.Info("some thing happend")
	// panic("errrr")
	//if err eventbus.send补偿 
	return &Response{A:1},errors.New("error")
}
//compensation 补偿怎么写


