package grpc

import (
	"context"
	"crud/pkg/controller"
	"crud/pkg/model"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type server struct {
	UnimplementedUserapiServer
}

var (
	port = flag.Int("port", 8201, "The server port")
)

func (s *server) DeleteUser(c context.Context, in *UserRequest) (*UserReply, error) {

	u, err := controller.DeleteUserCassandra(int(in.GetId()))
	if err != nil {
		return &UserReply{Message: "Error while deleteing", Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
	}
	return &UserReply{Message: "Deleted succesfully", Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
}

func (s *server) PutUser(c context.Context, in *UserPut) (*UserReply, error) {

	u := model.User{Id: int(in.GetId()), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: int(in.GetAge())}

	err := controller.UpdateUserCassandra(int(in.GetId()), u)

	if err != nil {
		return &UserReply{Message: "Error while updating", Id: in.GetId()}, nil
	}

	return &UserReply{Id: in.GetId(), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: in.GetAge()}, nil
}
func (s *server) PostUser(c context.Context, in *UserPost) (*UserReply, error) {
	u := model.User{Id: int(in.GetId()), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: int(in.GetAge())}
	err := controller.SaveUserCassandra(u)

	if err != nil {
		return &UserReply{Message: "Error while creating user", Id: in.GetId()}, nil
	}

	return &UserReply{Id: in.GetId(), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: in.GetAge()}, err

}

func (s *server) GetUsers(in *UserRequest, stream Userapi_GetUsersServer) error {
	// users := []model.User{}
	users, _ := controller.GetUsersCassandra()

	for _, u := range users {

		if err := stream.Send(&UserReply{Id: int64(u.Id), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}); err != nil {
			return err
		}
	}

	return nil
}
func (s *server) GetUser(c context.Context, in *UserRequest) (*UserReply, error) {

	u, err := controller.GetUserByIdCassandra(int(in.GetId()))

	if err != nil {
		return &UserReply{Message: "Error while getting user", Id: in.GetId()}, nil
	}

	return &UserReply{Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
}

func CreateGRPCServer() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterUserapiServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
