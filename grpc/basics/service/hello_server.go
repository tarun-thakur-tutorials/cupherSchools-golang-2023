package service

import (
	"context"
	"errors"
	"hello_grpc/proto"
)

type HelloServer struct {
	go_hello.UnimplementedMyServiceServer
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

/*
	HelloWorld(context.Context, *HelloMessage) (*HelloMessage, error)
	HelloNeeraj(context.Context, *HelloMessage) (*HelloMessage, error)
	HelloName(context.Context, *HelloMessage) (*HelloMessage, error)
	mustEmbedUnimplementedMyServiceServer()

*/

func (s *HelloServer) HelloWorld(
	ctx context.Context,
	req *go_hello.HelloMessage,
) (res *go_hello.HelloMessage, err error) {
	res = &go_hello.HelloMessage{}
	res.Body = "hello world"
	return
}

func (s *HelloServer) HelloName(
	ctx context.Context,
	req *go_hello.HelloMessage,
) (res *go_hello.HelloMessage, err error) {
	name := req.GetBody()

	if len(name) == 0 {
		err = errors.New("empty request body")
		return
	}

	res.Body = "hello " + name
	return
}

func (s *HelloServer) HelloNeeraj(
	ctx context.Context,
	req *go_hello.HelloMessage,
) (res *go_hello.HelloMessage, err error) {
	res.Body = "hello Neeraj"
	return
}

func (s *HelloServer) mustEmbedUnimplementedMyServiceServer() {}
