package date_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/cneill/date"
)

func TestStringJSON(t *testing.T) { //nolint:funlen // No sense splitting this up.
	t.Parallel()

	now := time.Date(2024, time.January, 1, 1, 0, 0, 1, time.UTC)

	type testStruct struct {
		ANSIC       date.String[date.ANSIC]       `json:"ansic"`
		DateOnly    date.String[date.DateOnly]    `json:"date_only"`
		DateTime    date.String[date.DateTime]    `json:"date_time"`
		Kitchen     date.String[date.Kitchen]     `json:"kitchen"`
		RFC1123     date.String[date.RFC1123]     `json:"rfc_1123"`
		RFC1123Z    date.String[date.RFC1123Z]    `json:"rfc_1123z"`
		RFC3339     date.String[date.RFC3339]     `json:"rfc_3339"`
		RFC3339Nano date.String[date.RFC3339Nano] `json:"rfc_3339_nano"`
		RFC822      date.String[date.RFC822]      `json:"rfc_822"`
		RFC822Z     date.String[date.RFC822Z]     `json:"rfc_822z"`
		RFC850      date.String[date.RFC850]      `json:"rfc_850"`
		RubyDate    date.String[date.RubyDate]    `json:"ruby_date"`
		Stamp       date.String[date.Stamp]       `json:"stamp"`
		StampMicro  date.String[date.StampMicro]  `json:"stamp_micro"`
		StampMilli  date.String[date.StampMilli]  `json:"stamp_milli"`
		StampNano   date.String[date.StampNano]   `json:"stamp_nano"`
		TimeOnly    date.String[date.TimeOnly]    `json:"time_only"`
		UnixDate    date.String[date.UnixDate]    `json:"unix_date"`
	}

	input := testStruct{
		ANSIC:       date.StringANSIC(now),
		DateOnly:    date.StringDateOnly(now),
		DateTime:    date.StringDateTime(now),
		Kitchen:     date.StringKitchen(now),
		RFC1123:     date.StringRFC1123(now),
		RFC1123Z:    date.StringRFC1123Z(now),
		RFC3339:     date.StringRFC3339(now),
		RFC3339Nano: date.StringRFC3339Nano(now),
		RFC822:      date.StringRFC822(now),
		RFC822Z:     date.StringRFC822Z(now),
		RFC850:      date.StringRFC850(now),
		RubyDate:    date.StringRubyDate(now),
		Stamp:       date.StringStamp(now),
		StampMicro:  date.StringStampMicro(now),
		StampMilli:  date.StringStampMilli(now),
		StampNano:   date.StringStampNano(now),
		TimeOnly:    date.StringTimeOnly(now),
		UnixDate:    date.StringUnixDate(now),
	}

	expected := `{
  "ansic": "Mon Jan  1 01:00:00 2024",
  "date_only": "2024-01-01",
  "date_time": "2024-01-01 01:00:00",
  "kitchen": "1:00AM",
  "rfc_1123": "Mon, 01 Jan 2024 01:00:00 UTC",
  "rfc_1123z": "Mon, 01 Jan 2024 01:00:00 +0000",
  "rfc_3339": "2024-01-01T01:00:00Z",
  "rfc_3339_nano": "2024-01-01T01:00:00.000000001Z",
  "rfc_822": "01 Jan 24 01:00 UTC",
  "rfc_822z": "01 Jan 24 01:00 +0000",
  "rfc_850": "Monday, 01-Jan-24 01:00:00 UTC",
  "ruby_date": "Mon Jan 01 01:00:00 +0000 2024",
  "stamp": "Jan  1 01:00:00",
  "stamp_micro": "Jan  1 01:00:00.000000",
  "stamp_milli": "Jan  1 01:00:00.000",
  "stamp_nano": "Jan  1 01:00:00.000000001",
  "time_only": "01:00:00",
  "unix_date": "Mon Jan  1 01:00:00 UTC 2024"
}`

	testBytes, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal test struct into JSON: %v", err)
	}

	if testString := string(testBytes); testString != expected {
		t.Errorf("got invalid JSON:\n%s\n", testString)
	}

	output := testStruct{}
	if err := json.Unmarshal(testBytes, &output); err != nil {
		t.Fatalf("failed to unmarshal marshaled JSON into test struct: %v", err)
	}

	testBytes, err = json.MarshalIndent(input, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal after unmarshalling JSON into test struct: %v", err)
	}

	if testString := string(testBytes); testString != expected {
		t.Errorf("got invalid JSON:\n%s\n", testString)
	}
}
