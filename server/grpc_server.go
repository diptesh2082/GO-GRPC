package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/diptesh2082/billing-software/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


type server struct {
    pb.UnimplementedGreeterServer
    // pb.UnimplementedStockServiceServer
}
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

// StreamStockPrices streams stock prices to the client
func (s *server) StreamStockPrices(req *pb.StockRequest, stream pb.Greeter_StreamStockPricesServer) error {
    log.Printf("Received request for stock: %s", req.StockSymbol)
    rand.Seed(time.Now().UnixNano())

    // Simulate streaming stock prices
    for i := 0; i < 10; i++ {
        price := rand.Float32()*100 + 50 // Random price between 50 and 150
        timestamp := time.Now().Format(time.RFC3339)

        res := &pb.StockResponse{
            StockSymbol: req.StockSymbol,
            Price:       price,
            Timestamp:   timestamp,
        }

        if err := stream.Send(res); err != nil {
            return err
        }

        time.Sleep(1 * time.Second) // Simulate delay between updates
    }

    return nil
}

func (s *server) StreamStockPricesClient(stream pb.Greeter_StreamStockPricesClientServer) error {
    log.Println("Receiving stock prices from client...")
    var count int
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            response := &pb.StockResponseT{
                Message: "Received all stock prices",
            }
            return stream.SendAndClose(response)
        }
        if err != nil {
            log.Printf("Error receiving stream: %v", err)
            return err
        }

        log.Printf("Received: %s with price: %.2f", req.StockSymbol, req.Price)
        count++
    }
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	// Register reflection service for debugging
	reflection.Register(grpcServer)

	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}