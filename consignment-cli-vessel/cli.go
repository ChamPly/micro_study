package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service-vessel/proto/consignment"
	"time"

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
	service := micro.NewService(
		micro.Name("go.micro.cli.consignment"),
	)

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", service.Client())

	service.Init()

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

		time.Sleep(time.Second * 10)
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
