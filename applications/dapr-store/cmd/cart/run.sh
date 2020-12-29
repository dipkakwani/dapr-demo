#!/bin/bash
dapr run --app-id cart --app-port 9001 --log-level debug realize start
#dapr run --app-id cart --app-port 9001 --log-level debug go run main.go routes.go
