package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service-micro/proto/consignment"

	micro "github.com/micro/go-micro"
)

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	// cmd.Init()

	service := micro.NewService(
		micro.Name("go.micro.cli.consignment"),
	)
	service.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())

	//conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	//if err != nil {
	//log.Fatalf("connect error:%+v", err)
	//}
	//defer conn.Close()

	//client := pb.NewShippingServiceClient(conn)
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error:%+v", err)
	}
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error:%+v", err)
	}

	log.Printf("created:%t", resp.Created)

	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignerts: %+v", err)
	}
	for _, c := range resp.Consignments {
		log.Printf("%+v", c)
	}
}
