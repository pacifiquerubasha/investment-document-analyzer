<template>
  <div class="results-container">
    <div class="results-header">
      <h2>Analysis Results</h2>
      <div class="results-counter" v-if="documents.length > 0">
        {{ documents.length }} document{{
          documents.length !== 1 ? "s" : ""
        }}
        analyzed
      </div>
    </div>

    <div v-if="documents.length === 0" class="no-results">
      <div class="empty-state-icon">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="48"
          height="48"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="1"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path
            d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"
          ></path>
          <polyline points="14 2 14 8 20 8"></polyline>
          <line x1="16" y1="13" x2="8" y2="13"></line>
          <line x1="16" y1="17" x2="8" y2="17"></line>
          <polyline points="10 9 9 9 8 9"></polyline>
        </svg>
      </div>
      <p>No documents analyzed yet.</p>
      <p class="empty-state-helper">Upload a PDF to get started.</p>
    </div>

    <div v-else class="results-wrapper">
      <div class="results-table">
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
              <td >
                <div class="file-icon">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <path
                      d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"
                    ></path>
                    <polyline points="14 2 14 8 20 8"></polyline>
                  </svg>
                </div>
                {{ doc.filename }}
              </td>
              <td>
                <div class="topic-tags">
                  <span
                    v-for="topic in doc.analysis.keyTopics"
                    :key="topic"
                    class="tag topic-tag"
                  >
                    {{ topic }}
                  </span>
                </div>
              </td>
              <td>
                <span
                  :class="[
                    'sentiment-badge',
                    getSentimentClass(doc.analysis.sentiment),
                  ]"
                >
                  {{ doc.analysis.sentiment }}
                </span>
              </td>
              <td>
                <div class="company-tags">
                  <span
                    v-for="company in doc.analysis.mentionedCompanies"
                    :key="company"
                    class="tag company-tag"
                  >
                    {{ company }}
                  </span>
                  <span
                    v-if="doc.analysis.mentionedCompanies.length === 0"
                    class="no-data"
                    >None</span
                  >
                </div>
              </td>
              <td class="date-cell">
                {{ formatDate(doc.analysis.analyzedAt) }}
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import { getSentimentClass, formatDate } from "../utils/formatters.js";

export default {
  name: "AnalysisResults",
  props: {
    documents: {
      type: Array,
      required: true,
    },
  },
  computed: {
    sortedDocuments() {
      return [...this.documents].sort(
        (a, b) =>
          new Date(b.analysis.analyzedAt) - new Date(a.analysis.analyzedAt)
      );
    },
  },
  methods: {
    getSentimentClass,
    formatDate,
  },
};
</script>

<style>
@import "../assets/styles/analysisResults.css";
</style>
