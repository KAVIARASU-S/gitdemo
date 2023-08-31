package main

import (
	//controller "Project1/Netxd_customer_controller"
	service "Project1/Netxd_dal_services"

	"log"
	"net"
	"Project1/Netxd_dal_model"

	"google.golang.org/grpc"
	//"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	customerService := &service.CustomerService{}
	netxddalmodel.RegisterBankingServiceServer(server, customerService)

	log.Println("Starting server...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
