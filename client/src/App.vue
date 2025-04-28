<template>
  <div class="container">
    <header>
      <h1>Investment Document Analyzer</h1>
      <p class="subtitle">Upload PDF documents for AI-simulated analysis</p>
    </header>
    
    <main>
      <DocumentUpload @document-analyzed="loadDocuments" />
      <AnalysisResults :documents="documents" />
    </main>
    
    <footer>
      <p>Mock Investment Research Tool - gRPC + Vue.js + Go Demo</p>
    </footer>
  </div>
</template>

<script>
import DocumentUpload from './components/DocumentUpload.vue'
import AnalysisResults from './components/AnalysisResults.vue'
import { DocumentAnalysisServiceClient } from './proto/document_grpc_web_pb.js'
import { ListDocumentsRequest } from './proto/document_pb.js'

export default {
  name: 'App',
  components: {
    DocumentUpload,
    AnalysisResults
  },
  data() {
    return {
      documents: [],
      client: null
    }
  },
  created() {
    this.client = new DocumentAnalysisServiceClient('http://localhost:8080')
    this.loadDocuments()
  },
  methods: {
    loadDocuments() {
      const request = new ListDocumentsRequest()
      
      this.client.listDocuments(request, {}, (err, response) => {
        if (err) {
          console.error('Error loading documents:', err)
          alert('Failed to load documents. Please ensure the server is running.')
          return
        }
        
        this.documents = response.getDocumentsList().map(doc => ({
          id: doc.getDocumentId(),
          filename: doc.getFilename(),
          analysis: {
            keyTopics: doc.getAnalysis().getKeyTopicsList(),
            sentiment: doc.getAnalysis().getSentiment(),
            mentionedCompanies: doc.getAnalysis().getMentionedCompaniesList(),
            analyzedAt: doc.getAnalysis().getAnalyzedAt()
          },
          uploadedAt: doc.getUploadedAt()
        }))
      })
    }
  }
}
</script>

<style>
@import './assets/styles/variables.css';
@import './assets/styles/app.css';
</style>