package sitemap

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Frequency type used for the freq field
type Frequency string

// Valid values for freq field
const (
	FrequencyAlways  = "always"
	FrequencyHourly  = "hourly"
	FrequencyDaily   = "daily"
	FrequencyWeekly  = "weekly"
	FrequencyMonthly = "monthly"
	FrequencyYearly  = "yearly"
	FrequencyNever   = "never"
)

func (f *Frequency) String() string {
	return string(*f)
}

// UnmarshalXML implements the xml.Unmarshaler interface
func (f *Frequency) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	v := new(string)

	err := d.DecodeElement(v, &start)
	if err != nil {
		return fmt.Errorf("failed to decode Frequency: %w", err)
	}

	*v = strings.TrimSpace(*v)

	if *v != FrequencyAlways &&
		*v != FrequencyHourly &&
		*v != FrequencyDaily &&
		*v != FrequencyWeekly &&
		*v != FrequencyMonthly &&
		*v != FrequencyYearly &&
		*v != FrequencyNever {

		return fmt.Errorf("invalid value for freq: %s", *v)
	}

	*f = *(*Frequency)(v)

	return nil
}
