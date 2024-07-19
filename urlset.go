package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
)

var (
	ErrorInvalidFormat = errors.New("invalid format")
)

type URLSet []URL

// UnmarshalURLSet unmarshals sitemaps (XMLs that starts/end with the `urlset` tag) and checks the required loc attribute.
//
// To unmarshal Sitemap Index, use the UnmarshalIndex.
func UnmarshalURLSet(data []byte) (URLSet, error) {

	if data == nil {
		return nil, fmt.Errorf("data is nil")
	}

	v := struct {
		URLs []URL `xml:"url"`
	}{}

	err := xml.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	for i := range v.URLs {
		if len(v.URLs[i].Loc.String()) == 0 {
			return nil, fmt.Errorf("loc is missing")
		}
	}

	return v.URLs, err
}

func (us URLSet) URLs() []URL {
	return []URL(us)
}
