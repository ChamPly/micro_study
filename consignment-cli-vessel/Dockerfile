FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

ADD consignment.json /app/consignment.json
ADD consignment-cli-vessel /app/consignment-cli-vessel

CMD ["./consignment-cli-vessel"]
