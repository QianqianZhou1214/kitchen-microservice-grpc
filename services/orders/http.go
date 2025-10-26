package main

import (
	"log"
	"net/http"

	"github.com/QianqianZhou1214/kitchen-microservice-grpc/services/orders/handler"
	"github.com/QianqianZhou1214/kitchen-microservice-grpc/services/orders/service"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)
	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
