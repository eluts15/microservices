build:
		protoc --go_out=plugins=grpc:. *.proto
		docker build -t user-cli .

run:
		docker run -e MICRO_REGISTRY=mdns consignment-cli
