package eventhandler

// import (
// 	"encoding/json"
// 	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
// 	"github.com/lj19950508/ddd-demo-go/pkg/logger"
// 	"github.com/streadway/amqp"
// )


// type UserEventHandler struct {
// 	logger   logger.Interface
// }

// func NewUserEventHandler(logger logger.Interface) *UserEventHandler {
// 	return &UserEventHandler{
// 		logger:logger,
// 	}
// }

// func (t *UserEventHandler) Dispatcher() eventbus.DispatchInfos {
// 	return eventbus.DispatchInfos{
// 		//默认使用user吧
// 		{EventName:"UserCreateEvent", Handle: t.Handler1},
		
// 	}
// }




// func (s *UserEventHandler) Handler1(d *amqp.Delivery)(any,error){
// 	req :=Response{}
// 	json.Unmarshal(d.Body,&req)

// 	s.logger.Info("some thing happend %v",req)
// 	return &Response{A:1},nil
// }


