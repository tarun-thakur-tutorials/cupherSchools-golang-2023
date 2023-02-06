package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	pcbook "pcbook/proto"
	"time"

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

func (server *LaptopServer) UploadImage(stream pcbook.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		log.Println(err)
		return err
	}

	laptopId := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()

	log.Println("got the request with image type: ", imageType, "and laptop id: ", laptopId)

	imageData := bytes.Buffer{}

	for {
		log.Println("waiting to recieve more data")
		time.Sleep(2 * time.Second)
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("no more data to recieve")
			break
		}

		if err != nil {
			return status.Errorf(codes.Unknown, "cannot recieve the chunk of data: %v", err)
		}

		chunk := req.GetChunkData()

		_, err = imageData.Write(chunk)

		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data")
		}
	}

	res := &pcbook.UploadImageResponse{Id: laptopId}

	err = stream.SendAndClose(res)

	if err != nil {
		return status.Error(codes.Unknown, err.Error())
	}

	fmt.Println("the data is recieved and response is sent")

	return nil
}
