package eventbusimpl

import (
	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
	"github.com/lj19950508/ddd-demo-go/pkg/rmq_rpc/client"
	"github.com/lj19950508/ddd-demo-go/pkg/rmq_rpc/server"
	"github.com/streadway/amqp"
)

type MqRpcEventBus struct{
	client *client.Client
	server *server.Server
}

func NewMqRpcEventBus() (*MqRpcEventBus,error){
	//创建client
	//创建爱你server
	client,err:=client.New("amqp://guest:guest@localhost:5672/","server","client")
	if(err!=nil){
		return nil,err
	}
	server,err:=server.New("amqp://guest:guest@localhost:5672/","server")
	if(err!=nil){
		return nil,err
	}
	// client *client.Client,server *server.Server
	return &MqRpcEventBus{
		client: client,
		server: server,
	},nil
}

//utilbus.Publish
func (s *MqRpcEventBus) Publish(evt *eventbus.Event) error {
	return s.client.RemoteCall(evt.Name,evt.Payload,&evt.Response)
}

//Onstart Subs....
func (s *MqRpcEventBus) Subscribe(name string, handler any) error {
	
	hd,ok:=handler.(func(*amqp.Delivery) (interface{}, error))
	if(ok){
		s.server.Subscribe(name,hd)
	}
	return nil
}

func (s *MqRpcEventBus) Start()  {
	s.server.GoConsumer()
}

func (s *MqRpcEventBus) Close(){
	s.client.Shutdown()
	s.server.Shutdown()

}