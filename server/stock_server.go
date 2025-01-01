package main

import (
    "log"
    "math/rand"
    "time"

    pb "github.com/diptesh2082/billing-software/proto"
)

// StockServiceServer implements the StockService gRPC interface
// type StockServiceServer struct {
//     pb.UnimplementedStockServiceServer
// }
// type server struct {
//     pb.UnimplementedGreeterServer
//     // pb.UnimplementedStockServiceServer
// }
// StreamStockPrices streams stock prices to the client
func (s *server) tStreamStockPrices(req *pb.StockRequest, stream pb.Greeter_StreamStockPricesServer) error {
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
