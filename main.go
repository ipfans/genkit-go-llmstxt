package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/mackee/go-readability"
)

// regexp to match URLs in the content, e.g., ""/go/docs/get-started-go/""
const httpRegexp = "/go/docs[^\"]+"

func parseContent(body []byte) (*readability.ReadabilityArticle, error) {
	options := readability.DefaultOptions()
	article, err := readability.Extract(string(body), options)
	if err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}
	return &article, nil
}

func main() {
	r, err := http.Get("https://genkit.dev")
	if err != nil {
		panic(fmt.Errorf("failed to get URL: %w", err))
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		panic(fmt.Errorf("failed to fetch URL: %s", r.Status))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(fmt.Errorf("failed to read response body: %w", err))
	}

	re := regexp.MustCompile(httpRegexp)
	matches := re.FindAll(body, -1)
	reader, writer := io.Pipe()

	go func() {
		for _, match := range matches {
			var (
				r       *http.Response
				b       []byte
				article *readability.ReadabilityArticle
			)
			r, err = http.Get("https://genkit.dev" + string(match))
			if err != nil {
				panic(fmt.Errorf("failed to get page: %w", err))
			}
			defer r.Body.Close()
			b, err = io.ReadAll(r.Body)
			if err != nil {
				panic(fmt.Errorf("failed to read page: %w", err))
			}
			article, err = parseContent(b)
			if err != nil {
				panic(fmt.Errorf("failed to parse content: %w", err))
			}
			_, _ = fmt.Fprintln(writer, readability.ToMarkdown(article.Root))
			_, _ = fmt.Fprintln(writer, "---")

		}
		_ = writer.Close()
	}()

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to copy: %v", err)
		os.Exit(1)
	}
}
