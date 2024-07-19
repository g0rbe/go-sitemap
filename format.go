package sitemap

import (
	"encoding/xml"
	"fmt"
)

const (
	FormatURLSet = "urlset"
	FormatIndex  = "sitemapindex"
)

// GetFormat unmarshals the XML data and returns the sitemap format.
// The format is either FormatURLSet or FormatIndex.
//
// If invalid format found, returns ErrorInvalidFormat
func GetFormat(data []byte) (string, error) {

	if data == nil {
		return "", fmt.Errorf("data is nil")
	}

	v := struct {
		XMLName xml.Name
	}{}

	err := xml.Unmarshal(data, &v)
	if err != nil {
		return "", err
	}

	if v.XMLName.Local != FormatURLSet && v.XMLName.Local != FormatIndex {
		return "", ErrorInvalidFormat
	}

	return v.XMLName.Local, err
}

// FetchFormat downloads the XML from the url, unmarshals it and returns the sitemap format and the downloaded data.
// The format is either FormatURLSet or FormatIndex.
//
// If invalid format found, returns ErrorInvalidFormat
func FetchFormat(url string) (string, []byte, error) {

	data, err := download(url)
	if err != nil {
		return "", nil, err
	}

	f, err := GetFormat(data)

	return f, data, err
}
