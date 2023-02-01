package main

import (
	"context"
	"flag"
	"fmt"
	go_hello "hello_grpc/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()

	log.Printf("dial server: %s \n", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial the server ", err)
	}

	client := go_hello.NewMyServiceClient(conn)

	req := &go_hello.HelloMessage{
		Body: "",
	}

	res, err := client.HelloWorld(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Body)
}
