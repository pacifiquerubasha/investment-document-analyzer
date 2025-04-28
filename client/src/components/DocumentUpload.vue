<template>
  <div class="upload-container">
    <h2>Upload Document</h2>
    <form @submit.prevent="handleSubmit" class="upload-form">
      <div class="file-input-wrapper">
        <input
          type="file"
          ref="fileInput"
          @change="handleFileChange"
          accept=".pdf"
          class="file-input"
          id="fileInput"
        />
        <label for="fileInput" class="file-label">
          <div class="file-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
              <polyline points="14 2 14 8 20 8"></polyline>
              <line x1="12" y1="12" x2="12" y2="18"></line>
              <line x1="9" y1="15" x2="15" y2="15"></line>
            </svg>
          </div>
          <div class="file-label-text">
            <span v-if="!selectedFile">Choose PDF file</span>
            <span v-else class="selected-filename">{{ selectedFile.name }}</span>
          </div>
        </label>
      </div>

      <button
        type="submit"
        :disabled="!selectedFile || isUploading"
        class="submit-button"
      >
        <span v-if="isUploading" class="loading-spinner"></span>
        <span>{{ isUploading ? "Analyzing..." : "Upload & Analyze" }}</span>
      </button>
    </form>

    <div v-if="error" class="error-message">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="10"></circle>
        <line x1="12" y1="8" x2="12" y2="12"></line>
        <line x1="12" y1="16" x2="12.01" y2="16"></line>
      </svg>
      {{ error }}
    </div>
  </div>
</template>

<script>
import { DocumentAnalysisServiceClient } from "../proto/document_grpc_web_pb.js";
import { AnalyzeDocumentRequest } from "../proto/document_pb.js";
import { readFileAsArrayBuffer } from "../utils/fileUtils.js";

export default {
  name: "DocumentUpload",
  emits: ["document-analyzed"],
  data() {
    return {
      selectedFile: null,
      isUploading: false,
      error: null,
      client: null,
    };
  },
  created() {
    this.client = new DocumentAnalysisServiceClient("http://localhost:8080");
  },
  methods: {
    handleFileChange(event) {
      const file = event.target.files[0];
      if (file && file.type === "application/pdf") {
        this.selectedFile = file;
        this.error = null;
      } else {
        this.selectedFile = null;
        this.error = "Please select a PDF file";
      }
    },

    async handleSubmit() {
      if (!this.selectedFile) return;

      this.isUploading = true;
      this.error = null;

      try {
        const arrayBuffer = await readFileAsArrayBuffer(this.selectedFile);
        const uint8Array = new Uint8Array(arrayBuffer);

        const request = new AnalyzeDocumentRequest();
        request.setFilename(this.selectedFile.name);
        request.setContent(uint8Array);

        // eslint-disable-next-line
        this.client.analyzeDocument(request, {}, (err, _response) => {
          if (err) {
            console.error("Error uploading document:", err);
            this.error = "Failed to analyze document. Please try again.";
            this.isUploading = false;
            return;
          }

          // Clear form
          this.selectedFile = null;
          this.$refs.fileInput.value = "";

          // Notify parent component
          this.$emit("document-analyzed");
          this.isUploading = false;
        });
      } catch (error) {
        console.error("Error reading file:", error);
        this.error = "Failed to read document. Please try again.";
        this.isUploading = false;
      }
    },
  },
};
</script>

<style>
@import '../assets/styles/documentUpload.css';
</style>