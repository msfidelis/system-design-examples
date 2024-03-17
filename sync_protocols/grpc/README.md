```bash
brew install grpc protobuf protoc-gen-go-grpc protoc-gen-go
```


```bash
protoc --go_out=../server/ --go-grpc_out=../server/ imc.proto
protoc --go_out=../client/ --go-grpc_out=../client/ imc.proto
```