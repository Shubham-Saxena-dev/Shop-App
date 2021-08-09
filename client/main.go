package main

import (
	"Shop-app/client/services"
	"Shop-app/db"
	"Shop-app/proto/stubs/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const PORT = 8081

func main() {
	db.InitializeDatabase()
	initializeServer()
}

func initializeServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Unable to open server at port {}", PORT)
		panic(err)
	}

	server := grpc.NewServer()
	order.RegisterOrderServiceServer(server, &services.OrderService{})
	reflection.Register(server)

	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
