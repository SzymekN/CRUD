package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "https://github.com/SzymekN/CRUD/tree/main/pkg/grpc"
)

type server struct {
	pb.UnimplementedUserapiServer
}

func (s *server) GetUser(c echo.Context, in *pb.UserRequest) (*pb.UserReply, error){

}

func CreateServer(){
	s := grpc.NewServer()

}