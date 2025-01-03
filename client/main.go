package main

import (

	"log"
	// "github.com/diptesh2082/billing-software/client"
	pb "github.com/diptesh2082/billing-software/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new StockService client
	client := pb.NewGreeterClient(conn)

	// Call the function to stream stock prices
	// CallStreamStockPricesClient(client)
	CallStreamStockBiServer(client)

}
