package eventhandler

import (
	"github.com/grafana/grafana/pkg/bus"
	"github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
)

type EventHandler struct {
	eventBus bus.Bus
	logger   logger.Interface
}

func NewEventHandler(eventBus bus.Bus, logger logger.Interface) *EventHandler {
	eventBus.AddEventListener(func(evt *user.EvtUserCreate) error {
		//TODO eventbus 需要改版 可以 同步或异步，同步可以返回值， 目前这个grafa的方案比较一般 试试改成mq。。 或者自己写 go fun
		logger.Info("收到了eventbus2的消息")
		return nil
	})
	eventBus.AddEventListener(func(evt *user.EvtUserCreate) error {
		logger.Info("收到了eventbus1的消息")
		return nil
	})
	return &EventHandler{
		eventBus,
		logger,
	}
}
