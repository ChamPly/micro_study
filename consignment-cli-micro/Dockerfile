FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

ADD consignment.json /app/consignment.json
ADD consignment-cli-micro /app/consignment-cli-micro

CMD ["./consignment-cli-micro"]
