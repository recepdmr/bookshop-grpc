package main

import (
	"bookshop/client/bookshop/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Code removed for brevity

	client := pb.NewInventoryClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	bookList, _ := client.GetBookList(context.Background(), &pb.GetBookListRequest{})

	bookDetail, _ := client.GetBookDetailById(context.Background(), &pb.GetBookDetailByIdRequest{
		Id: 2,
	})

	log.Printf("book list: %v \n", bookList)

	log.Printf("book detail: %v\n", bookDetail)

}
