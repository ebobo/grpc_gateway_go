package main

import (
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

var opt struct {
	GRPCAddr string `long:"grpc-addr" default:":9092" description:"gRPC listen address"`
}

func main() {
	_, err := flags.ParseArgs(&opt, os.Args)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	var userManagementServer = NewUserServer()

	if err := userManagementServer.Run(opt.GRPCAddr); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
