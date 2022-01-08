package main

import (
	"context"
	"log"
	"time"

	"github.com/ebobo/grpc_gateway_go/pkg/go/userpb/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	userServerAddress = "localhost:9092"
)

func main() {
	conn, err := grpc.Dial(userServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}
	defer conn.Close()
	c := userpb.NewUserServiceClient(conn)

	// https://go.dev/blog/context
	// In Go servers, each incoming request is handled in its own goroutine.
	// Request handlers often start additional goroutines to access backends such as databases
	// and RPC services. The set of goroutines working on a request typically needs access to
	// request-specific values such as the identity of the end user, authorization tokens,
	// and the request’s deadline. When a request is canceled or times out, all the goroutines
	// working on that request should exit quickly so the system can reclaim any resources they are using.
	// At Google, we developed a context package that makes it easy to pass request-scoped values,
	// cancelation signals, and deadlines across API boundaries to all the goroutines involved
	// in handling a request. The package is publicly available as context.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	var newUsers = make(map[string]int32)

	newUsers["Espen"] = 25
	newUsers["Qi"] = 39
	newUsers["Stig"] = 50

	for name, age := range newUsers {
		r, err := c.CreateUser(ctx, &userpb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user %v", err)
		}
		log.Printf(`User Details: Name: %s Age: %d Id: %d`, r.GetName(), r.GetAge(), r.GetId())
	}

	params := &userpb.GetUsersParams{}
	r, err := c.GetUser(ctx, params)
	if err != nil {
		log.Fatalf("could not get user list %v", err)
	}
	log.Print("\n User List: \n")
	log.Printf(" %v\n", r.GetUsers())

}

// use go mod tidy to download all the pakages we imported
