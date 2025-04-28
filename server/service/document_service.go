package service

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	pb "investment-document-analyzer/server/proto"
	"investment-document-analyzer/server/models"
)

type DocumentService struct {
	pb.UnimplementedDocumentAnalysisServiceServer
	documents map[string]*models.Document
	mu        sync.RWMutex
}

func NewDocumentService() *DocumentService {
	return &DocumentService{
		documents: make(map[string]*models.Document),
	}
}

// AnalyzeDocument handles the document upload and analysis request
func (s *DocumentService) AnalyzeDocument(ctx context.Context, req *pb.AnalyzeDocumentRequest) (*pb.AnalyzeDocumentResponse, error) {
	// Generate a unique document ID
	docID := uuid.New().String()
	
	// Simulate analysis by generating random data
	analysis := generateMockAnalysis()
	
	// Create and store the document
	doc := &models.Document{
		ID:         docID,
		Filename:   req.GetFilename(),
		Content:    req.GetContent(),
		Analysis:   analysis,
		UploadedAt: time.Now(),
	}
	
	s.mu.Lock()
	s.documents[docID] = doc
	s.mu.Unlock()
	
	return &pb.AnalyzeDocumentResponse{
		DocumentId: docID,
		Filename:   req.GetFilename(),
		Analysis:   analysis,
	}, nil
}

// ListDocuments returns a list of all analyzed documents
func (s *DocumentService) ListDocuments(ctx context.Context, req *pb.ListDocumentsRequest) (*pb.ListDocumentsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	var documents []*pb.DocumentSummary
	for _, doc := range s.documents {
		documents = append(documents, &pb.DocumentSummary{
			DocumentId: doc.ID,
			Filename:   doc.Filename,
			Analysis:   doc.Analysis,
			UploadedAt: doc.UploadedAt.Format(time.RFC3339),
		})
	}
	
	return &pb.ListDocumentsResponse{
		Documents: documents,
	}, nil
}

// generateMockAnalysis creates fake analysis results for demonstration
func generateMockAnalysis() *pb.AnalysisResult {
	// Predefined data for mock analysis
	topics := []string{
		"AI Technology Trends",
		"Market Volatility",
		"ESG Investments",
		"Cryptocurrency Analysis",
		"Global Economic Outlook",
		"Interest Rate Predictions",
		"Tech Sector Growth",
		"Renewable Energy Opportunities",
		"Emerging Markets",
		"Supply Chain Resilience",
	}
	
	sentiments := []string{
		"Positive",
		"Neutral",
		"Negative",
		"Mixed",
	}
	
	companies := []string{
		"AAPL", "MSFT", "GOOGL", "AMZN", "TSLA",
		"NVDA", "META", "NFLX", "JPM", "BAC",
		"V", "MA", "PG", "JNJ", "KO",
		"DIS", "INTC", "AMD", "PYPL", "CRM",
	}
	
	// Randomly select analysis components
	rand.Seed(time.Now().UnixNano())
	
	selectedTopics := selectRandom(topics, 2+rand.Intn(3))
	selectedSentiment := sentiments[rand.Intn(len(sentiments))]
	selectedCompanies := selectRandom(companies, 3+rand.Intn(5))
	
	return &pb.AnalysisResult{
		KeyTopics:          selectedTopics,
		Sentiment:          selectedSentiment,
		MentionedCompanies: selectedCompanies,
		AnalyzedAt:         time.Now().Format(time.RFC3339),
	}
}

// selectRandom selects n random items from a slice without duplicates
func selectRandom(slice []string, n int) []string {
	if n > len(slice) {
		n = len(slice)
	}
	
	// Fisher-Yates shuffle
	shuffled := make([]string, len(slice))
	copy(shuffled, slice)
	
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}
	
	return shuffled[:n]
}