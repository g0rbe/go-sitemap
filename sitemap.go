package sitemap

import "fmt"

type Sitemap []URL

// Fetch downloads the Sitemap XML from the given URL and returns a slice of URLs.
// If the sitemap on the given URL is an index, recursively fetch the the URL the get the URLs.
func Fetch(url string) (Sitemap, error) {

	data, err := download(url)
	if err != nil {
		return nil, err
	}

	f, err := GetFormat(data)
	if err != nil {
		return nil, fmt.Errorf("format error: %w", err)
	}

	switch f {
	case FormatURLSet:

		us, err := UnmarshalURLSet(data)
		return Sitemap(us), err

	case FormatIndex:

		urls, err := UnmarshalIndex(data)
		if err != nil {
			return nil, err
		}

		var si Sitemap

		for i := range urls {

			s, err := Fetch(urls[i].String())
			if err != nil {
				return nil, err
			}

			si = append(si, s...)

		}

		return si, nil

	default:

		return nil, fmt.Errorf("invalid format: %s", f)

	}
}

// // Unmarshal check the Sitemap format and unmarshal the XML with the appropriate function.
// func Unmarshal(data []byte) (Sitemap, error) {

// 	if data == nil {
// 		return nil, fmt.Errorf("data is nil")
// 	}

// 	f, err := GetFormat(data)
// 	if err != nil {
// 		return nil, fmt.Errorf("format error: %w", err)
// 	}

// 	switch f {
// 	case FormatURLSet:
// 		return UnmarshalURLSet(data)
// 	case FormatIndex:
// 		return UnmarshalIndex(data)
// 	default:
// 		return nil, fmt.Errorf("invalid format: %s", f)
// 	}
// }
