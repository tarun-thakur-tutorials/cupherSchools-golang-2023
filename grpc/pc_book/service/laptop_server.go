package service

import (
	"context"
	"fmt"
	"log"
	pcbook "pcbook/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore // im-memory laptop store
	pcbook.UnimplementedLaptopServiceServer
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{
		Store: LaptopStore{
			data: map[string]*pcbook.Laptop{},
		},
	}
}

/*
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	mustEmbedUnimplementedLaptopServiceServer()
*/

func (l *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pcbook.CreateLaptopRequest,
) (*pcbook.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("recieved the create-laptop request with the id: %s\n", laptop.Id)

	if len(laptop.GetId()) == 0 {
		return nil, status.Error(codes.Internal, "invalid id")
	}

	id := laptop.GetId()

	if _, ok := l.Store.data[id]; ok {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("laptop already exists with  id %v", id),
		)
	}

	l.Store.data[id] = laptop

	res := &pcbook.CreateLaptopResponse{
		Id: id,
	}

	return res, nil
}
func (l *LaptopServer) mustEmbedUnimplementedLaptopServiceServer() {}
