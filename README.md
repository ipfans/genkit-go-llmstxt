# genkit-go-llmstxt

A Go tool that scrapes the [Firebase Genkit Go](https://genkit.dev) documentation and converts it to [llms-txt](https://llmstxt.org/).

## Overview

This tool fetches the main Genkit documentation page, extracts all documentation URLs, and converts each page to clean llms-txt. The output is aggregated into a single file that can be used as context for Large Language Models. (~47k tokens)

## Features

- Fetches content from https://genkit.dev
- Automatically discovers documentation URLs using regex pattern matching
- Converts HTML pages to clean llms-txt using readability extraction
- Outputs consolidated documentation to a single file
- Concurrent processing for efficient scraping

## Installation

```bash
go mod download
```

## Usage

### Using Task (recommended)

```bash
task
```

### Using Go directly

```bash
go run main.go > llms-full.txt
```

The tool will:

1. Fetch the main Genkit documentation page
2. Extract all documentation URLs matching the pattern `/go/docs/*`
3. Process each page to extract readable content
4. Output markdown-formatted content to stdout (or redirect to file)

## Dependencies

- [go-readability](https://github.com/mackee/go-readability) - For HTML content extraction and markdown conversion

## License

MIT License - see [LICENSE](LICENSE) file for details.
