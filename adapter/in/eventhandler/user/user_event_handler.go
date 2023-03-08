package eventhandler

import (
	"github.com/grafana/grafana/pkg/bus"
	"github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
)

//这个改成类似api的处理 
//然后有一个eventbushandler 是
//用户子域的事件处理

//EventHandler inteface has a router  routerInfo [message,func]
//TODO eventbus need eventhandler[]  and eventbus foreach .AddEventListenr()
//   Queue=>func (evt *user.EvtUserCreate) error
// so just need eventbu

type UserEventHandler struct {
	eventBus bus.Bus
	logger   logger.Interface
}

func NewUserEventHandler(eventBus bus.Bus, logger logger.Interface) *UserEventHandler {
	eventBus.AddEventListener(func(evt *user.EvtUserCreate) error {
		logger.Info("收到了eventbus2的消息")
		return nil
	})
	eventBus.AddEventListener(func(evt *user.EvtUserCreate) error {
		logger.Info("收到了eventbus1的消息")
		return nil
	})
	return &UserEventHandler{
		eventBus,
		logger,
	}
}
