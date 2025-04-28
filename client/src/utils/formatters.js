/**
 * Get the CSS class name for sentiment badges
 * @param {string} sentiment - The sentiment value
 * @returns {string} - CSS class name
 */
export function getSentimentClass(sentiment) {
    const sentimentLower = sentiment.toLowerCase();
    switch (sentimentLower) {
      case 'positive':
        return 'positive';
      case 'negative':
        return 'negative';
      case 'mixed':
        return 'mixed';
      default:
        return 'neutral';
    }
  }
  
  /**
   * Format a date string to a localized format
   * @param {string} dateString - The date string to format
   * @returns {string} - Formatted date string
   */
  export function formatDate(dateString) {
    return new Date(dateString).toLocaleString(undefined, {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  }