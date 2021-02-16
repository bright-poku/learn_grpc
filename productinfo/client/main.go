package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "productinfo/client/ecommerce"
)

  const (
	  address = "localhost:50051"
  )

  func main() {
	  //setup a connection to the Server
	  conn, err := grpc.Dial(address, grpc.WithInsecure())
	  if err != nil {
		  log.Fatalf("did not connect to: %v", err)
	  }
	  defer conn.Close()
	  c := pb.NewProductInfoClient(conn)

	  name := "Apple Iphone 12"
	  description := "meet the new Apple Iphone"
	  //price := float32(1000.0)
	  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	  defer cancel()

	  r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Id: "1"})
	  if err != nil {
		  log.Fatalf("Could not add product: %v", err)
	  }
	  log.Printf("Product ID: %s added successfully", r.Value)

	  product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	  if err != nil {
		  log.Fatalf("Could not get product: %v", err)
	  }
	  log.Printf("Product: ", product.String())
  }