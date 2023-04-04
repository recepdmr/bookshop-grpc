# Bookshop GRPC services implementation

That repo contains grpc service implementation which used common proto files on different programming languages

## Supported languages
  - Go
  - .NET
  
## How to use these services through grpcurl

[grpcurl](https://github.com/fullstorydev/grpcurl) is a command-line tool that lets you interact with gRPC servers. It's basically curl for gRPC servers. 

###  How to curl call to GetBookDetailById service method

```sh
    grpcurl -plaintext -d '{"id":32}' localhost:8080 Inventory/GetBookDetailById
```