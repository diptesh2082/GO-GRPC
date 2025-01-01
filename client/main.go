package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	CallStreamStockPricesClient(client)

}
func CallStreamStockPricesClient(client pb.GreeterClient) {
	stream, err := client.StreamStockPricesClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	stocks := []struct {
		symbol string
		price  float32
	}{
		{"AAPL", 150.25},
		{"GOOGL", 2800.50},
		{"AMZN", 3450.75},
		{"MSFT", 299.99},
	}

	for _, stock := range stocks {
		log.Printf("Sending stock: %s with price: %.2f", stock.symbol, stock.price)
		if err := stream.Send(&pb.StockRequestT{
			StockSymbol: stock.symbol,
			Price:       stock.price,
		}); err != nil {
			log.Fatalf("Failed to send stock: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	log.Printf("Server response: %s", response.Message)
}

func CallStreamStockPrices(client pb.GreeterClient) {
	// Request to stream stock prices
	fmt.Println("------------------------------")
	req := &pb.StockRequest{StockSymbol: "AAPL"}
	stream, err := client.StreamStockPrices(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling StreamStockPrices: %v", err)
	}

	// Receive stock prices from the stream
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Printf("Stream ended: %v", err)
			break
		}
		log.Printf("Received stock price: %s - $%.2f at %s", res.StockSymbol, res.Price, res.Timestamp)
	}
	time.Sleep(2 * time.Second)
}
