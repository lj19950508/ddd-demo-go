package command

import (
	"context"
	"flag"
	"log"
	"time"
	userpkg "github.com/lj19950508/ddd-demo-go/domain/user"
	"github.com/lj19950508/ddd-demo-go/pkg/eventbus"
	pb "github.com/lj19950508/ddd-demo-go/protos/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserCommandService interface {
	Create(cmd *CreateCommand) error
	Update(cmd *UpdateCommand) error
	Delete(id int64) error
}

//---------------------------

type UserCommandImpl struct {
	userRepository userpkg.UserRepository
	eventBus       eventbus.EventBus
}

func NewUserCommandImpl(userRepository userpkg.UserRepository, eventBus eventbus.EventBus) UserCommandService {
	return &UserCommandImpl{
		userRepository: userRepository,
		eventBus:       eventBus,
	}
}

func (t UserCommandImpl) Create(cmd *CreateCommand) error {
	user := userpkg.NewUser(0, cmd.Name)
	if err := t.userRepository.Add(user); err != nil {
		return err
	}
	addr := flag.String("addr", "127.0.0.1:8081", "the address to connect to")
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c:=pb.NewUserCenterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err:=c.Login(ctx,&pb.SaveEvent{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())

	// res:=userpkg.EventResUserCreate{}
	// event:=userpkg.NewEventUserCreate(1,&res)
	// evt:=eventbus.NewEvent(event.Eventname,event)
	// err := t.eventBus.Publish(evt)
	// if err != nil {
		// return err
	// }/
	// fmt.Println(res)

	return nil
}

func (t UserCommandImpl) Update(cmd *UpdateCommand) error {
	user, err := t.userRepository.Load(cmd.ID)
	if err != nil {
		return err
	}
	user.Name = cmd.Name
	if err = t.userRepository.Save(user); err != nil {
		return err
	}
	return nil
}

func (t UserCommandImpl) Delete(id int64) error {
	user, err := t.userRepository.Load(id)
	if err != nil {
		return err
	}

	if err := t.userRepository.Remove(user); err != nil {
		return err
	}
	return nil
}
