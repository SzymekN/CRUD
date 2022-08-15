package grpc

import (
	"context"
	"crud/pkg/model"
	"crud/pkg/storage"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/gocql/gocql"
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
	cas := storage.GetCassandraInstance()

	u := model.User{}

	if err := cas.Query(`Select id, firstname, lastname, age from userapi.users where id=?`, in.GetId()).Consistency(gocql.One).Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age); err != nil {
		fmt.Println(err)
	}

	if err := cas.Query(`Delete from userapi.users where id=?`, in.GetId()).Exec(); err != nil {
		fmt.Println(err)
	}
	return &UserReply{Id: in.GetId(), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}, nil
}

func (s *server) PutUser(c context.Context, in *UserPut) (*UserReply, error) {
	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Select id, firstname, lastname, age from userapi.users where id=?`, in.GetId()).Consistency(gocql.One).Scan(in.GetId(), in.GetFirstname(), in.GetLastname(), in.GetAge()); err != nil {
		fmt.Println(err)
	}

	if err := cas.Query(`Update userapi.users set firstname=?, lastname=?, age=? where id=?`, in.GetFirstname(), in.GetLastname(), in.GetAge(), in.GetId()).Exec(); err != nil {
		fmt.Println(err)
	}
	return &UserReply{Id: in.GetId(), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: in.GetAge()}, nil
}
func (s *server) PostUser(c context.Context, in *UserPost) (*UserReply, error) {
	cas := storage.GetCassandraInstance()
	if err := cas.Query(`Insert into userapi.users(id, firstname,lastname, age) values (?,?,?,?)`, in.GetId(), in.GetFirstname(), in.GetLastname(), in.GetAge()).Exec(); err != nil {
		fmt.Println(err)
	}
	return &UserReply{Id: in.GetId(), Firstname: in.GetFirstname(), Lastname: in.GetLastname(), Age: in.GetAge()}, nil

}

func (s *server) GetUsers(in *UserRequest, stream Userapi_GetUsersServer) error {
	// users := []model.User{}
	u := model.User{}

	cas := storage.GetCassandraInstance()

	iter := cas.Query(`Select id, firstname, lastname, age from userapi.users`).Iter()
	for iter.Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age) {
		// stream.Send(&UserReply{Id: int64(u.Id), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)})
		// u = model.User{}
		if err := stream.Send(&UserReply{Id: int64(u.Id), Firstname: u.Firstname, Lastname: u.Lastname, Age: int64(u.Age)}); err != nil {
			return err
		}
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return nil
}
func (s *server) GetUser(c context.Context, in *UserRequest) (*UserReply, error) {

	id := in.GetId()
	cas := storage.GetCassandraInstance()
	u := model.User{}

	if err := cas.Query(`Select id, firstname, lastname, age from userapi.users where id=?`, id).Consistency(gocql.One).Scan(&u.Id, &u.Firstname, &u.Lastname, &u.Age); err != nil {
		fmt.Println(err)
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
