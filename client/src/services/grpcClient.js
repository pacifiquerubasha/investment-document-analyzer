// grpcClient.js - Wrapper for gRPC Web client
import { 
    DocumentAnalysisServiceClient, 
    AnalyzeDocumentRequest, 
    ListDocumentsRequest 
  } from '../proto/index.js'
  
  // Create a single client instance
  const client = new DocumentAnalysisServiceClient('http://localhost:8080')
  
  export function analyzeDocument(filename, content) {
    return new Promise((resolve, reject) => {
      const request = new AnalyzeDocumentRequest()
      request.setFilename(filename)
      request.setContent(content)
      
      client.analyzeDocument(request, {}, (err, response) => {
        if (err) {
          reject(err)
        } else {
          resolve(response)
        }
      })
    })
  }
  
  export function listDocuments() {
    return new Promise((resolve, reject) => {
      const request = new ListDocumentsRequest()
      
      client.listDocuments(request, {}, (err, response) => {
        if (err) {
          reject(err)
        } else {
          resolve(response)
        }
      })
    })
  }
  
  export default {
    analyzeDocument,
    listDocuments
  }