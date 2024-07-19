package sitemap

import (
	"encoding/xml"
	"fmt"
)

type Index []URL

// UnmarshalIndex unmarshals Sitemap Indexes (XMLs that starts/end with the `sitemapindex` tag) and checks the required loc attribute.
//
// To unmarshal Sitemap, use the UnmarshalURLSet.
func UnmarshalIndex(data []byte) (Index, error) {

	if data == nil {
		return nil, fmt.Errorf("data is nil")
	}

	v := struct {
		URLs []URL `xml:"sitemap"`
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
