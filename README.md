# Testing Guide for Investment Document Analyzer

This guide provides step-by-step instructions for testing the Investment Document Analyzer application.

## Prerequisites

Before you begin, ensure you have installed:

- Go 1.24+
- Node.js and npm
- Protocol Buffers compiler (protoc)
- Docker (for Envoy proxy)
- [Protocol Buffer plugins](#installing-protocol-buffer-plugins)

## Step-by-Step Testing Process

### 1. Setup the Project

```bash
# Setup the project (generates proto files and installs dependencies)
make setup
```

### 2. Build the Backend

```bash
make build-server
```

### 3. Test Components Individually

#### 3.1 Test the Backend

```bash
make run-server
```

You should see output like:

```
Starting Go server on port 50051...
Server listening at [::]:50051
```

Keep this terminal open and open a new terminal for the next step.

#### 3.2 Test the Envoy Proxy

```bash
make run-proxy
```

You should see Docker starting the Envoy container:

```
Starting Envoy proxy...
```

Check if the container is running:

```bash
docker ps | grep envoy
```

#### 3.3 Test the Frontend

```bash
make run-frontend
```

You should see output indicating that the Vue.js development server is running, usually on port 3000.

### 4. Test All Components Together

If all individual components work, you can run them all together:

```bash
# Stop any running services first
make stop

# Run all services
make run-all
```

### 5. Use the Application

1. Open your browser and navigate to [http://localhost:3000](http://localhost:3000)
2. You should see the Investment Document Analyzer interface
3. Use the file upload to select a PDF document
4. Click "Upload & Analyze" to process the document
5. The analysis results should appear in the table below

## Expected Results

After uploading a document, you should see:

- The document name in the results table
- Randomly generated key topics
- A sentiment value (Positive, Neutral, Negative, or Mixed)
- Randomly selected company ticker symbols
- Timestamp of when the analysis was performed

### 6. Cleaning Up

When you're done testing, stop all services:

```bash
make stop
```

If you want to remove build artifacts:

```bash
make clean
```

## Troubleshooting

### Common Issues

#### Protocol Buffer Plugin Errors

If you see errors like:

```bash
protoc-gen-go: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go_out: protoc-gen-go: Plugin failed with status code 1
```

Make sure you have:

1. Installed the protocol buffer plugins:

```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Added your Go bin directory to your PATH:

- Linux/macOS:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

- Windows:

```bash
set PATH=%PATH%;%USERPROFILE%\go\bin
```

3. Restarted your terminal after making these changes and run make setup again
