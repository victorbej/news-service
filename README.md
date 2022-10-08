1. Install gRPC and Protobuf on your pc (Guide - https://grpc.io/docs/languages/go/quickstart/)
2. Git clone third party google api https://fuchsia.googlesource.com/third_party/googleapis/
3. For go-gen type "protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     pkg/api/news.proto
   " from root project directory
4. Server port: 8000
5. Client port: 8080