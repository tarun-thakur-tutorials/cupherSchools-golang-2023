package main

import (
	"context"
	"flag"
	"log"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("created the laptop with id: %v\n", res.Id)
}
