package models

import (
	"time"

	pb "investment-document-analyzer/server/proto"
)

// Document represents a document with its analysis results
type Document struct {
	ID         string
	Filename   string
	Content    []byte
	Analysis   *pb.AnalysisResult
	UploadedAt time.Time
}