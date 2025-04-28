<template>
  <div class="results-container">
    <h2>Analysis Results</h2>
    
    <div v-if="documents.length === 0" class="no-results">
      No documents analyzed yet. Upload a PDF to get started.
    </div>
    
    <div v-else class="results-table">
      <table>
        <thead>
          <tr>
            <th>Filename</th>
            <th>Key Topics</th>
            <th>Sentiment</th>
            <th>Mentioned Companies</th>
            <th>Analyzed At</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="doc in sortedDocuments" :key="doc.id">
            <td>{{ doc.filename }}</td>
            <td>
              <div class="topic-tags">
                <span v-for="topic in doc.analysis.keyTopics" :key="topic" class="tag topic-tag">
                  {{ topic }}
                </span>
              </div>
            </td>
            <td>
              <span :class="['sentiment-badge', getSentimentClass(doc.analysis.sentiment)]">
                {{ doc.analysis.sentiment }}
              </span>
            </td>
            <td>
              <div class="company-tags">
                <span v-for="company in doc.analysis.mentionedCompanies" :key="company" class="tag company-tag">
                  {{ company }}
                </span>
              </div>
            </td>
            <td>{{ formatDate(doc.analysis.analyzedAt) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AnalysisResults',
  props: {
    documents: {
      type: Array,
      required: true
    }
  },
  computed: {
    sortedDocuments() {
      return [...this.documents].sort((a, b) => 
        new Date(b.analysis.analyzedAt) - new Date(a.analysis.analyzedAt)
      )
    }
  },
  methods: {
    getSentimentClass(sentiment) {
      const sentimentLower = sentiment.toLowerCase()
      switch (sentimentLower) {
        case 'positive':
          return 'positive'
        case 'negative':
          return 'negative'
        case 'mixed':
          return 'mixed'
        default:
          return 'neutral'
      }
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleString()
    }
  }
}
</script>

<style scoped>
.results-container {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
  margin-top: 0;
  color: var(--primary-color);
}

.no-results {
  text-align: center;
  padding: 40px;
  color: #666;
  background-color: var(--background-color);
  border-radius: 4px;
}

.results-table {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th, td {
  padding: 12px;
  text-align: left;
  border-bottom: 1px solid var(--border-color);
}

th {
  background-color: var(--background-color);
  font-weight: 600;
  color: var(--primary-color);
}

tr:hover {
  background-color: #f8f9fa;
}

.tag {
  display: inline-block;
  padding: 4px 8px;
  margin: 2px;
  border-radius: 4px;
  font-size: 0.85em;
}

.topic-tags, .company-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.topic-tag {
  background-color: #e8f4fd;
  color: #0366d6;
}

.company-tag {
  background-color: #f0f0f0;
  color: #333;
}

.sentiment-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.9em;
  font-weight: 500;
}

.sentiment-badge.positive {
  background-color: #d4edda;
  color: #155724;
}

.sentiment-badge.neutral {
  background-color: #fff3cd;
  color: #856404;
}

.sentiment-badge.negative {
  background-color: #f8d7da;
  color: #721c24;
}

.sentiment-badge.mixed {
  background-color: #e2e3e5;
  color: #383d41;
}
</style>