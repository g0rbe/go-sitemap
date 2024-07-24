package sitemap

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"
)

type ContentTypeError struct {
	ContentType string
}

func NewContentTypeError(ct string) ContentTypeError {
	return ContentTypeError{ContentType: ct}
}

func (cte ContentTypeError) Error() string {
	return "invalid content type: \"" + cte.ContentType + "\""
}

var (
	ErrTooManyRequests = errors.New("429 Too Many Requests")
	ErrBadRequest      = errors.New("400 Bad Request")
	ErrNotFound        = errors.New("404 Not Found")
	ErrUnauthorized    = errors.New("401 Unauthorized")
	ErrForbidden       = errors.New("403 Forbidden")
)

func download(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/xml, text/xml")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: %w", err)
	}
	defer resp.Body.Close()

	// Check response code
	switch resp.StatusCode {
	case 200:
		break
	case 400:
		return nil, ErrBadRequest
	case 401:
		return nil, ErrUnauthorized
	case 403:
		return nil, ErrForbidden
	case 404:
		return nil, ErrNotFound
	case 429:
		return nil, ErrTooManyRequests
	default:
		return nil, fmt.Errorf(resp.Status)
	}

	// Check Content-Type
	mediaType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse media type: %w", err)
	}

	switch true {
	case mediaType == "application/xml",
		mediaType == "text/xml",
		mediaType == "application/atom+xml":

		return io.ReadAll(resp.Body)

	case mediaType == "application/x-gzip",
		mediaType == "application/octet-stream" && strings.HasSuffix(url, ".gz"):

		data, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("gzip reader error: %w", err)
		}
		defer data.Close()

		return io.ReadAll(data)

	default:

		return nil, NewContentTypeError(mediaType)
	}

}
