.PHONY: all help setup generate-proto build-server run-server run-client run-proxy run-all clean stop status

BACKEND_DIR := server
FRONTEND_DIR := client
PROXY_DIR := proxy
PROTO_DIR := proto

# Default target
all: setup build-server

help:
	@echo "Available commands:"
	@echo "  make setup           - Install dependencies and generate proto files"
	@echo "  make generate-proto  - Generate protobuf files for Go and JS"
	@echo "  make build-server    - Build the Go server"
	@echo "  make run-server      - Run only the Go server"
	@echo "  make run-client      - Run only the Vue.js frontend"
	@echo "  make run-proxy       - Run only the Envoy proxy"
	@echo "  make run-all         - Run all services (server, client, proxy)"
	@echo "  make stop            - Stop all running services"
	@echo "  make clean           - Clean build artifacts"

# Check and install dependencies
setup:
	@echo "Installing Go protoc plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "Installing protoc-gen-grpc-web..."
	npm install -g protoc-gen-grpc-web
	@echo "Setting up client dependencies..."
	cd $(FRONTEND_DIR) && npm install
	@echo "Setting up server dependencies..."
	cd $(BACKEND_DIR) && go mod tidy
	@echo "Generating protobuf files..."
	$(MAKE) generate-proto

# Generate protobuf code for both Go and JS
generate-proto:
	# Make sure directories exist
	mkdir -p $(PROTO_DIR)
	mkdir -p $(BACKEND_DIR)/proto
	mkdir -p $(FRONTEND_DIR)/src/proto
	
	# Generate Go code
	protoc -I$(PROTO_DIR) \
		--go_out=$(BACKEND_DIR)/proto --go_opt=paths=source_relative \
		--go-grpc_out=$(BACKEND_DIR)/proto --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/document.proto
	
	# Generate JS code
	protoc -I$(PROTO_DIR) \
		--js_out=import_style=commonjs:$(FRONTEND_DIR)/src/proto \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:$(FRONTEND_DIR)/src/proto \
		$(PROTO_DIR)/document.proto

# Build the backend executable
build-server:
	@echo "Building Go server..."
	mkdir -p $(BACKEND_DIR)/bin
	cd $(BACKEND_DIR) && go build -o bin/server main.go

# Run the backend server
run-server:
	@echo "Starting Go server on port 50051..."
	cd $(BACKEND_DIR) && ./bin/server

# Run the frontend
run-client:
	@echo "Starting Vue.js frontend..."
	cd $(FRONTEND_DIR) && npm run serve

# Run Envoy proxy
run-proxy:
	@echo "Starting Envoy proxy..."
	cd $(PROXY_DIR) && docker-compose up -d

# Run all services
run-all:
	@echo "Starting all services..."
	@$(MAKE) -j3 run-server run-proxy run-client

# Stop all running services
stop:
	@echo "Stopping services..."
	-pkill -f "$(BACKEND_DIR)/bin/server" || true
	-cd $(PROXY_DIR) && docker-compose down || true
	-pkill -f "$(FRONTEND_DIR)" || true
	@echo "All services stopped."

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BACKEND_DIR)/bin/
	rm -rf $(BACKEND_DIR)/proto/*.pb.go
	rm -rf $(FRONTEND_DIR)/src/proto/*_pb.js
	rm -rf $(FRONTEND_DIR)/node_modules
	rm -rf $(FRONTEND_DIR)/dist
	@echo "Cleanup complete!"