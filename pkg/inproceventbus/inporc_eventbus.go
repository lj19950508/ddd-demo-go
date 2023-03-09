package inproceventbus

import (
	"errors"
	"fmt"

	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
)

//同步进程内的eventbus
type InProcEventBus struct {
	//维持一个队列
	handlers map[string][]eventbus.EventHandler //可行
	//return handler
	results map[int64]chan any //有用 没做超时
	//存一个map  这个event类型 ，以及对应的补偿 （也是一个event）
	events []*eventbus.Event //可行

	errorExcute []*eventbus.Event //补偿机制谁来决定 ？ 客户端,客户端 可以告诉接收端 错误发生了 应该发送哪个队列补偿让客户端去做补偿
}

//是个同步进程列队列
func (s *InProcEventBus) Publish(evt *eventbus.Event) error {
	//客户端请求 ， 订阅执行失败则返回错误 预期之外的错误不用处理 交给 httphandle
	// s.grafanaBus.Publish()
	if s.handlers == nil {
		//
		return errors.New("hanlder not exsits")
	}
	//事件入列
	s.events = append(s.events, evt)
	if evt.Return {
		s.results[evt.Id] = make(chan any)
		result := <-s.results[evt.Id]
		evt.ExcuteResult = result
	}

	return nil

}
func (s *InProcEventBus) Subscribe(name string, handler eventbus.EventHandler) error {
	s.handlers[name] = append(s.handlers[name], handler)	
	return nil
}

func (s *InProcEventBus) Consume(evt *eventbus.Event) {
	defer func() {
		if err := recover(); err != nil {
			//panic了
			s.errorExcute = append(s.errorExcute, evt)
			if evt.Return {
				s.results[evt.Id] <- nil
			}
		}
	}()
	for evtName, handlers := range s.handlers {
		if evtName == evt.Name {
			for _, handler := range handlers {
				handler(evt)
				if evt.Return {
					s.results[evt.Id] <- evt.ExcuteResult
				}
			}
		}
	}
}

func (s *InProcEventBus) Compensation(evt *eventbus.Event) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("补偿错误")
			//补偿失败 打印日志
		}
	}()
	comEvt := eventbus.NewEvent(0, evt.Compensation, evt.Payload)
	comEvt.ExcuteResult = false
	s.Publish(comEvt)
}

func (s *InProcEventBus) StartConsume() {
	go func() {
		for {
			if len(s.events) > 0 {
				evt := s.events[0]
				s.events = s.events[1:]
				go s.Consume(evt)
			}
		}
	}()
}

func (s *InProcEventBus) StartCompensation() {
	go func() {
		for {
			if len(s.errorExcute) > 0 {
				evt := s.errorExcute[0]
				s.errorExcute = s.errorExcute[1:]
				go s.Compensation(evt)
			}
		}
	}()

}

func NewInProcEventBus() *InProcEventBus {
	handlers := make(map[string][]eventbus.EventHandler)

	//return ha
	results := make(map[int64]chan any)
	errorExcute := make([]*eventbus.Event, 0)
	events := make([]*eventbus.Event, 0)
	return &InProcEventBus{
		handlers:    handlers,
		results:     results,
		errorExcute: errorExcute,
		events:      events,
	}
}
