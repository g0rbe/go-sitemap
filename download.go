package sitemap

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func download(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("accept", "application/xml, text/xml")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	// Check response code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	// Check Content-Type
	contentType, _, _ := strings.Cut(resp.Header.Get("content-type"), ";")

	switch true {
	case contentType == "application/xml", contentType == "text/xml", contentType == "application/atom+xml":

		return io.ReadAll(resp.Body)

	case contentType == "application/x-gzip", contentType == "application/octet-stream" && strings.HasSuffix(url, ".gz"):

		data, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("gzip reader error: %w", err)
		}
		defer data.Close()

		return io.ReadAll(data)

	default:

		return nil, fmt.Errorf("invalid content type: %s", resp.Header.Get("content-type"))
	}

}
