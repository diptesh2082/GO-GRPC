package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	// "time"

	pb "github.com/diptesh2082/billing-software/proto"
	// "google.golang.org/grpc"
)

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

func CallStreamStockPricesServer(client pb.GreeterClient) {
	// Request to stream stock prices
	fmt.Println("------------------------------")
	req := &pb.StockRequest{StockSymbol: "AAPL"}
	stream, err := client.StreamStockPricesServer(context.Background(), req)
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


func CallStreamStockBiServer(client pb.GreeterClient) {
	stream, err := client.StreamStockPricesBi(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2) // for both sender and receiver goroutines

	// Sending data to the server
	go func() {
		defer wg.Done()
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
				log.Printf("Failed to send stock: %v", err)
				return
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// Receiving data from the server
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server finished sending data.")
				return
			}
			if err != nil {
				log.Printf("Error receiving data: %v", err)
				return
			}
			log.Printf("Received from server: %s", res.Message)
		}
	}()

	// Wait for both goroutines to complete
	wg.Wait()
}
