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
          <span v-if="!selectedFile">Choose PDF file</span>
          <span v-else>{{ selectedFile.name }}</span>
        </label>
      </div>

      <button
        type="submit"
        :disabled="!selectedFile || isUploading"
        class="submit-button"
      >
        {{ isUploading ? "Analyzing..." : "Upload & Analyze" }}
      </button>
    </form>

    <div v-if="error" class="error-message">
      {{ error }}
    </div>
  </div>
</template>

<script>
import { DocumentAnalysisServiceClient } from "../proto/document_grpc_web_pb.js";
import { AnalyzeDocumentRequest } from "../proto/document_pb.js";

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
        const arrayBuffer = await this.readFileAsArrayBuffer(this.selectedFile);
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

    readFileAsArrayBuffer(file) {
      return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = (e) => resolve(e.target.result);
        reader.onerror = (e) => reject(e);
        reader.readAsArrayBuffer(file);
      });
    },
  },
};
</script>

<style scoped>
.upload-container {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
  margin-top: 0;
  color: var(--primary-color);
}

.upload-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.file-input-wrapper {
  position: relative;
}

.file-input {
  display: none;
}

.file-label {
  display: inline-block;
  padding: 12px 24px;
  background-color: var(--background-color);
  border: 2px dashed var(--border-color);
  border-radius: 4px;
  cursor: pointer;
  text-align: center;
  width: 100%;
  transition: all 0.3s ease;
}

.file-label:hover {
  border-color: var(--secondary-color);
  background-color: white;
}

.submit-button {
  padding: 12px 24px;
  background-color: var(--secondary-color);
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s ease;
}

.submit-button:hover:not(:disabled) {
  background-color: #2980b9;
}

.submit-button:disabled {
  background-color: #bdc3c7;
  cursor: not-allowed;
}

.error-message {
  color: var(--error-color);
  margin-top: 10px;
  padding: 10px;
  background-color: #ffeaea;
  border-radius: 4px;
}
</style>
