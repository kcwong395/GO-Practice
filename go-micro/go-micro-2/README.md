source: https://ewanvalentine.io/microservices-in-golang-part-2/

### Keywords
docker, go-micro

### Build docker image
sudo docker build -t shippy-service-consignment .

### Create a container from image
sudo docker run -p 50051:50051 shippy-service-consignment

### Update protobuf 
protoc --proto_path=. --go_out=. --micro_out=. proto/consignment/consignment.proto