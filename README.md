## Protobuf concepts:
- Message
- Service
- RPC methods
- Repeated

Takeaways: 
1. .proto file defines data and service interface
2. using protoc to generate code in any languages (here Golang)
3. implementing interface in server, calling RPC methods in client
4. gRPC serializing the requests into Protobuf binary and sending to server
5. server sends it back by Protobuf
6. faster than JSON, and type safety

orders.pb.go: Defining the strut of messages (data) in Go
orders_grpc.pg.go: Defining gRPC client/server interface
