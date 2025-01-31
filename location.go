package sitemap

import (
	"bytes"
	"net/url"
)

// Location is the URL of the page. This URL must begin with the protocol (such as http) and end with a trailing slash, if your web server requires it. This value must be less than 2,048 characters.
//
// Example:
//
//	<loc>http://www.example.com/</loc>
type Location []byte

func (l Location) Equal(loc Location) bool {
	if l == nil || loc == nil {
		return false
	}

	return bytes.Equal(l, loc)
}

func (l Location) URL() (*url.URL, error) {
	return url.Parse(l.String())
}

func (l Location) String() string {
	return string(l)
}
