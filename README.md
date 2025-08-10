# Text Search Engine using Go

This project is a simple full-text search engine implemented in Go. It demonstrates how to parse, index, and search large text datasets efficiently using basic information retrieval techniques.

---

## Features

- Loads and parses a compressed Wikipedia abstract XML dump.
- Tokenizes, lowercases, removes stopwords, and stems words for indexing.
- Builds an inverted index for fast full-text search.
- Supports command-line search queries.

---

## Implementation Details

### 1. Document Parsing

- Used Go's `encoding/xml` and `compress/gzip` to efficiently parse large, compressed XML files.
- Defined a `document` struct to map XML fields (`title`, `url`, `abstract`).

### 2. Text Processing Pipeline

- **Tokenization:** Split text into words using Unicode-aware rules.
- **Lowercasing:** Converted all tokens to lowercase for normalization.
- **Stopword Removal:** Filtered out common English stopwords.
- **Stemming:** Used the Snowball stemmer to reduce words to their root forms.

### 3. Indexing

- Built an inverted index (`map[string][]int`) mapping tokens to document IDs.
- Ensured efficient addition and lookup for fast search.

### 4. Search

- Processed queries through the same pipeline as documents.
- Used intersection of posting lists to find documents matching all query terms.

### 5. Command-Line Interface

- Used Go's `flag` package to accept the data file path and search query as arguments.

---

## How to Run

1. **Place the Wikipedia abstract dump** (e.g., `enwiki-latest-abstract1.xml.gz`) in the project directory.
2. **Build and run:**
   ```sh
   go run main.go -p enwiki-latest-abstract1.xml.gz -q "your search query"
   ```
   - `-p` specifies the path to the compressed XML dump.
   - `-q` specifies the search query.

---

## Things I Learned

- **Efficient File Handling:**  
  Processing large, compressed files without loading everything into memory at once.
- **Text Normalization:**  
  The importance of tokenization, stopword removal, and stemming in search engines.
- **Inverted Index:**  
  How an inverted index enables fast full-text search.
- **Go Language Features:**  
  Using Go's standard library for file I/O, XML parsing, and command-line argument parsing.
- **Performance Considerations:**  
  Optimizing preprocessing and indexing for speed and memory usage.
- **Git Large File Handling:**  
  Managing large files in version control and understanding GitHub's file size limits.

---

## Future Improvements

- Add support for phrase and fuzzy search.
- Implement ranking (e.g., TF-IDF).
- Build a web interface for interactive search.
- Use concurrency for faster indexing and searching.

---

**Author:** Raja Prem Sai
**License:** MIT