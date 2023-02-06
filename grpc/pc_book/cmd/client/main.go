package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	pcbook "pcbook/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")

	flag.Parse()
	log.Printf("dial server: %s\n", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	laptopClient := pcbook.NewLaptopServiceClient(conn)

	req := &pcbook.CreateLaptopRequest{
		Laptop: &pcbook.Laptop{
			Id:          "123424",
			Brand:       "Lenovo",
			Name:        "Yoga",
			Weight:      &pcbook.Laptop_WeightKg{WeightKg: 1},
			PriceUsd:    900,
			ReleaseYear: 2017,
			Cpu: &pcbook.CPU{
				Brand:         "Intel",
				NumberCores:   8,
				NumberThreads: 32,
			},
			Ram: &pcbook.Memory{
				Unit: pcbook.Unit_GIGABYTE,
				Val:  8,
			},
			Gpu: []*pcbook.GPU{},
			Storage: []*pcbook.Storage{
				{
					Driver: pcbook.Storage_SSD,
					Memory: &pcbook.Memory{Unit: pcbook.Unit_GIGABYTE, Val: 1024},
				},
			},
			Screen: &pcbook.Screen{
				Resolution: &pcbook.Resolution{
					Height: 16,
					Width:  16,
				},
				Panel: pcbook.Screen_IPS,
			},
			Keyboard: &pcbook.Keyboard{
				Backlit: true,
				Layout:  pcbook.Keyboard_QWERTY,
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("created the laptop with id: %v\n", res.Id)

	// file stream (upload)

	imgPath := "tmp/test.png"
	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatal("cannot open the image file: ", err)
	}
	defer file.Close()

	ctx, cancel = context.WithTimeout(context.Background(), 5000*time.Second)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)

	Imgreq := &pcbook.UploadImageRequest{
		Data: &pcbook.UploadImageRequest_Info{
			Info: &pcbook.ImageInfo{
				LaptopId:  "123421",
				ImageType: filepath.Ext(imgPath), // filepath is a library for path of the directory
			},
		},
	}
	err = stream.Send(Imgreq)
	if err != nil {
		log.Fatal("cannot send image info: ", err)
	}

	reader := bufio.NewReader(file)

	buffer := make([]byte, 1024) // size of the chunk = 1024 bytes = 1 kb

	for {

		// time.Sleep(2 * time.Second)
		n, err := reader.Read(buffer)

		if err == io.EOF {
			fmt.Println("got EOF")
			break
		}

		if err != nil {
			log.Fatal("cannot read the chunk to the buffer: ", err)
		}

		req := &pcbook.UploadImageRequest{
			Data: &pcbook.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		log.Println("sending data \n\n")
		time.Sleep(2 * time.Second)
		err = stream.Send(req)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("cannot sent stream ", err)
		}
	}

	Imgres, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot recieve response", err)
	}

	log.Printf("image is uploaded with id: %v, size: %v\n", Imgres.GetId(), Imgres.GetSize())
}
