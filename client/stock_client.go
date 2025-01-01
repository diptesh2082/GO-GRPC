package main

import (
	"context"
	"fmt"
	"log"
	"time"

	// "time"

	pb "github.com/diptesh2082/billing-software/proto"
	// "google.golang.org/grpc"
)

// Exported function to call stream stock prices
// func CallStreamStockPrices(client pb.GreeterClient) {
// 	// Request to stream stock prices
// 	fmt.Println("------------------------------")
// 	req := &pb.StockRequest{StockSymbol: "AAPL"}
// 	stream, err := client.StreamStockPrices(context.Background(), req)
// 	if err != nil {
// 		log.Fatalf("Error while calling StreamStockPrices: %v", err)
// 	}

// 	// Receive stock prices from the stream
// 	for {
// 		res, err := stream.Recv()
// 		if err != nil {
// 			log.Printf("Stream ended: %v", err)
// 			break
// 		}
// 		log.Printf("Received stock price: %s - $%.2f at %s", res.StockSymbol, res.Price, res.Timestamp)
// 	}
// 	time.Sleep(2 * time.Second)
// }

func CallSayHello(client pb.GreeterClient) {
	// Prepare the request
	req := &pb.HelloRequest{Name: "World"}

	// Call the SayHello method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Print the response
	fmt.Printf("Greeting: %s\n", resp.GetMessage())
}