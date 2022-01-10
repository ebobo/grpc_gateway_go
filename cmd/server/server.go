package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"github.com/ebobo/grpc_gateway_go/pkg/go/pb/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UserServer that implements the pb API
type UserServer struct {
	userList *pb.UserList
	pb.UnimplementedUserServiceServer
}

// NewUserServer creates a new Service instance
func NewUserServer() *UserServer {
	return &UserServer{userList: &pb.UserList{}}
}

// Run grpc server
func (server *UserServer) Run(port string) error {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("unable to listen %v", err)
	}
	gs := grpc.NewServer()
	reflection.Register(gs)

	pb.RegisterUserServiceServer(gs, server)

	log.Printf("server listening at %v", lis.Addr())
	return gs.Serve(lis)
}

// CreateUser implementation
func (server *UserServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Handle CreateUser %v", in.GetName())
	var userID = int32(rand.Intn(1000))

	createdUser := &pb.User{Id: userID, Name: in.GetName(), Age: in.GetAge(), Type: in.GetType()}
	server.userList.Users = append(server.userList.Users, createdUser)

	return createdUser, nil
}

// GetUser implementation
func (server *UserServer) GetUser(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return server.userList, nil
}
