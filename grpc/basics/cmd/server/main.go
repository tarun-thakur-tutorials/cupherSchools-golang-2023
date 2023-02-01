package main

import (
	"flag"
	"fmt"
	go_hello "hello_grpc/proto"
	"hello_grpc/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// flag.DataType("name of the flag", default_value, details)
	port := flag.Int("port", 0, "the server port")

	flag.Parse() // call it after defining all the flags

	log.Printf("start server on port: %d", *port)

	helloServer := service.NewHelloServer()

	grpcServer := grpc.NewServer()

	go_hello.RegisterMyServiceServer(grpcServer, helloServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Could not start the server: ", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Could not start the server: ", err)
	}
}
