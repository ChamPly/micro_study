build: 
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/shippy/vessel-service proto/vessel/vessel.proto
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run --rm --net=host -e MICRO_SERVER_ADDRESS=:50052 -e MICRO_REGISTER=mdns vessel-service
