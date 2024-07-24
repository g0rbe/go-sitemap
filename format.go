package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
)

type Format byte

const (
	FormatURLSet Format = 1 // "urlset"
	FormatIndex  Format = 2 // "sitemapindex"
)

var (
	ErrInvalidFormat      = errors.New("invalid format")
	ErrURLSetFormat       = errors.New("urlset format")
	ErrSitemapIndexFormat = errors.New("sitemapindex format")
)

func (f Format) String() string {
	switch f {
	case FormatURLSet:
		return "urlset"
	case FormatIndex:
		return "sitemapindex"
	default:
		return strconv.FormatUint(uint64(f), 10)
	}
}

// GetFormat unmarshals the XML data and returns the sitemap format.
// The format is either FormatURLSet or FormatIndex.
//
// If invalid format found, returns ErrorInvalidFormat
func GetFormat(data []byte) (Format, error) {

	if data == nil {
		return 0, fmt.Errorf("data is nil")
	}

	v := struct {
		XMLName xml.Name
	}{}

	err := xml.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}

	switch v.XMLName.Local {
	case "urlset":
		return FormatURLSet, nil
	case "sitemapindex":
		return FormatIndex, nil
	default:
		return 0, ErrInvalidFormat
	}
}

// FetchFormat downloads the XML from the url, unmarshals it and returns the sitemap format and the downloaded data.
// The format is either FormatURLSet or FormatIndex.
//
// If invalid format found, returns ErrorInvalidFormat
func FetchFormat(url string) (Format, []byte, error) {

	data, err := download(url)
	if err != nil {
		return 0, nil, err
	}

	f, err := GetFormat(data)

	return f, data, err
}
