# gRPC

1. Download proto buf from the link given below. (protoc-24.0-win64)

```
https://github.com/protocolbuffers/protobuf/releases/tag/v24.1
```

2. After downloading, extract it and paste its bin folder path in env.

3. After creating the proto file, excute the following command to create grpc.pb and pb file.
Proto_path will be your proto file path and go_out path will be where your file will be generated
```
 protoc --proto_path=proto proto/*.proto --go_out=gen/
```
```
 protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
```
4. First run your server and then client using following command go run ..