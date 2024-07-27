#!/bin/bash

services=("product" "user" "cart" "payment" "api-gateway")

for service in "${services[@]}"; do
    echo "Starting $service service..."
    cp .env.$service .env
    SERVICE_NAME=$service nohup go run /github.com/Nixa001/micromart/services/$service/cmd/main.go > /dev/null 2>&1 &
done

echo "All services started."
