BINARY=go-grpc-api-gateway

proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go

engine:
	go build -o ${BINARY} cmd/*.go

docker:
	docker build -t ${BINARY} .

docker_run:
	docker run --name ${BINARY} -p3000:3000 -d --network bridge-network ${BINARY}

docker_stop:
	docker stop ${BINARY}
	
docker_rm:
	docker rm $(shell docker ps -a -f name=${BINARY} -q)

.PHONY: proto server engine docker docker_run docker_stop docker_rm
