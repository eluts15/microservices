FROM golang:1.9.0 as builder

RUN go get -u github.com/eluts15/shipping/consignment-cli

WORKDIR /go/src/github.com/eluts15/shipping/consignment-cli

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM debian:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/eluts15/shipping/consignment-cli/consignment-cli .

ENTRYPOINT ["./consignment-cli", "consignment.json"]
CMD ["./consignment-cli"]
