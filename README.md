
# 📄 Investment Document Analyzer (Vue.js + Go + gRPC)

A simple research tool that allows users to upload PDF documents and receive AI-simulated analysis results. This project tests:

- gRPC + Go backend development
- Protobuf definition and code generation
- Vue.js frontend with gRPC-Web integration
- gRPC-Web proxy setup and communication flow

---

## 🎯 Objective

To simulate document analysis in a mock investment research tool using:
- **Golang with gRPC** for the backend
- **Protobufs** for API contracts
- **Vue.js** for the web frontend
- **Envoy or grpcwebproxy** for gRPC-Web communication

---

## 🛠 Tech Stack

- **Backend:** Golang with gRPC
- **Frontend:** Vue.js with gRPC-Web
- **Protobufs:** For request/response schemas
- **Proxy Layer:** Envoy or grpcwebproxy

---

## ✨ Functionality

### 1. 📤 Upload Investment Document
- Upload a PDF file via frontend
- Send the file over gRPC to the Go backend

### 2. 🧠 Simulate AI Document Analysis
Backend simulates analysis by returning:
- Key topics (e.g., “Tech Trends”)
- Sentiment (e.g., “Neutral”)
- Mentioned companies (e.g., “AAPL”, “TSLA”)

### 3. 📄 View Analysis Results
- Show insights on screen (topics, sentiment, companies)
- Maintain history of uploaded documents + results

---

## 📦 Protobuf API Definition (document.proto)

feel free to do your own proto definations here

## 🔧 Backend (Go)

- Implements `DocumentAnalysisService`
- Stores documents + analysis results in memory
- Randomly generates mock insights from predefined values
- Runs on `localhost:50051` for gRPC communication

---

## 🖥 Frontend (Vue.js)

- Upload PDF via a form
- Submit button to upload + analyze
- Display table of all uploaded docs + analysis results
- Integrates with backend via gRPC-Web


## ✅ Deliverables

- `proto/document.proto`
- Go-based `server/` with gRPC logic
- Vue.js `client/` with form + analysis viewer
- Envoy config for proxying gRPC-Web
- README with setup and usage docs

---

## 📋 Evaluation Criteria

- ✅ Protobuf definitions are correct and functional
- ✅ Backend is fully working with in-memory logic
- ✅ gRPC-Web correctly bridges Vue and Go
- ✅ Clean code with organized project structure
- ✅ Basic but intuitive UI

---

## 🙌 Notes

No real AI/ML models are required — just fake/mock data is fine for simulation.

