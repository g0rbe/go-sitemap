package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidChangeFrequency = errors.New("invalid changefreq")
)

// ChangeFrequency type used for the <changefreq> field
type ChangeFrequency byte

// Valid values for ChangeFrequency
var (
	ChangeFreqAlways  ChangeFrequency = 1 // "always"
	ChangeFreqHourly  ChangeFrequency = 2 // "hourly"
	ChangeFreqDaily   ChangeFrequency = 3 // "daily"
	ChangeFreqWeekly  ChangeFrequency = 4 // "weekly"
	ChangeFreqMonthly ChangeFrequency = 5 // "monthly"
	ChangeFreqYearly  ChangeFrequency = 6 // "yearly"
	ChangeFreqNever   ChangeFrequency = 7 // "never"
)

// ParseChangeFreq parses ChangeFrequency from v.
//
// If v is not a valid ChangeFrequency, returns ChangeFrequency(255)
func ParseChangeFreq(v string) ChangeFrequency {

	switch v {
	case "always":
		return ChangeFreqAlways
	case "hourly":
		return ChangeFreqHourly
	case "daily":
		return ChangeFreqDaily
	case "weekly":
		return ChangeFreqWeekly
	case "monthly":
		return ChangeFreqMonthly
	case "yearly":
		return ChangeFreqYearly
	case "never":
		return ChangeFreqNever
	default:
		return 255
	}
}

func (f ChangeFrequency) String() string {
	switch f {
	case ChangeFreqAlways:
		return "always"
	case ChangeFreqHourly:
		return "hourly"
	case ChangeFreqDaily:
		return "daily"
	case ChangeFreqWeekly:
		return "weekly"
	case ChangeFreqMonthly:
		return "monthly"
	case ChangeFreqYearly:
		return "yearly"
	case ChangeFreqNever:
		return "never"
	default:
		return strconv.Itoa(int(f))
	}
}

func (f ChangeFrequency) IsEmpty() bool {
	return f == 0
}

func (f ChangeFrequency) IsSet() bool {
	return f != 0
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (f *ChangeFrequency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	var v string

	err := d.DecodeElement(&v, &start)
	if err != nil {
		return fmt.Errorf("failed to decode ChangeFrequency: %w", err)
	}

	v = strings.TrimSpace(v)

	c := ParseChangeFreq(v)

	if c == 0 || c == 255 {
		return fmt.Errorf("%w: %v", ErrInvalidChangeFrequency, v)
	}

	*f = c

	return nil
}

func (f ChangeFrequency) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name.Local = "changefreq"

	return e.EncodeElement(f.String(), start)
}
