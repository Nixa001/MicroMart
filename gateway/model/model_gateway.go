package model

import (
	"log"
	"net/url"
)

type Gateway struct {
	ProductService string
	CartService    string
	PaymentService string
	UserService    string
}

func NewGateway() *Gateway {
	return &Gateway{
		ProductService: "http://localhost:8081",
		CartService:    "http://localhost:8082",
		PaymentService: "http://localhost:8083",
		UserService:    "http://localhost:8084",
	}
}

func (g *Gateway) ParseURL(urlStr string) *url.URL {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}
	return u
}
