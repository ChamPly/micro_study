build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .
run:
	docker run --rm --net=host -e MICRO_REGISTER=mdns consignment-cli
