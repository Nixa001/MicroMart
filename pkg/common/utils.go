package common

import (
	"log"
)

func GetServicePort(serviceName string) string {
	switch serviceName {
	case "product":
		return "8080"
	case "user":
		return "8081"
	case "cart":
		return "8082"
	case "payment":
		return "8083"
	case "api-gateway":
		return "8000"
	default:
		log.Fatalf("Unknown service: %s", serviceName)
		return ""
	}
}
