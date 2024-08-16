package date

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type StringFormat interface {
	Format() string
}

// ANSIC = Mon Jan _2 15:04:05 2006
type ANSIC string

func StringANSIC(input time.Time) String[ANSIC] { return String[ANSIC](input) }
func (a ANSIC) Format() string                  { return time.ANSIC }

// UnixDate = Mon Jan _2 15:04:05 MST 2006
type UnixDate string

func StringUnixDate(input time.Time) String[UnixDate] { return String[UnixDate](input) }
func (u UnixDate) Format() string                     { return time.UnixDate }

// RubyDate = Mon Jan 02 15:04:05 -0700 2006
type RubyDate string

func StringRubyDate(input time.Time) String[RubyDate] { return String[RubyDate](input) }
func (r RubyDate) Format() string                     { return time.RubyDate }

// RFC822 = 02 Jan 06 15:04 MST
type RFC822 string

func StringRFC822(input time.Time) String[RFC822] { return String[RFC822](input) }
func (r RFC822) Format() string                   { return time.RFC822 }

// RFC822Z = 02 Jan 06 15:04 -0700
type RFC822Z string

func StringRFC822Z(input time.Time) String[RFC822Z] { return String[RFC822Z](input) }
func (r RFC822Z) Format() string                    { return time.RFC822Z }

// RFC850 = Monday, 02-Jan-06 15:04:05 MST
type RFC850 string

func StringRFC850(input time.Time) String[RFC850] { return String[RFC850](input) }
func (r RFC850) Format() string                   { return time.RFC850 }

// RFC1123 = Mon, 02 Jan 2006 15:04:05 MST
type RFC1123 string

func StringRFC1123(input time.Time) String[RFC1123] { return String[RFC1123](input) }
func (r RFC1123) Format() string                    { return time.RFC1123 }

// RFC1123Z = Mon, 02 Jan 2006 15:04:05 -0700
type RFC1123Z string

func StringRFC1123Z(input time.Time) String[RFC1123Z] { return String[RFC1123Z](input) }
func (r RFC1123Z) Format() string                     { return time.RFC1123Z }

// RFC3339 = 2006-01-02T15:04:05Z07:00
type RFC3339 string

func StringRFC3339(input time.Time) String[RFC3339] { return String[RFC3339](input) }
func (r RFC3339) Format() string                    { return time.RFC3339 }

// RFC3339Nano = 2006-01-02T15:04:05.999999999Z07:00
type RFC3339Nano string

func StringRFC3339Nano(input time.Time) String[RFC3339Nano] { return String[RFC3339Nano](input) }
func (r RFC3339Nano) Format() string                        { return time.RFC3339Nano }

// Kitchen = 3:04PM
type Kitchen string

func StringKitchen(input time.Time) String[Kitchen] { return String[Kitchen](input) }
func (r Kitchen) Format() string                    { return time.Kitchen }

// Stamp = Jan _2 15:04:05
type Stamp string

func StringStamp(input time.Time) String[Stamp] { return String[Stamp](input) }
func (r Stamp) Format() string                  { return time.Stamp }

// StampMilli = Jan _2 15:04:05.000
type StampMilli string

func StringStampMilli(input time.Time) String[StampMilli] { return String[StampMilli](input) }
func (r StampMilli) Format() string                       { return time.StampMilli }

// StampMicro = Jan _2 15:04:05.000000
type StampMicro string

func StringStampMicro(input time.Time) String[StampMicro] { return String[StampMicro](input) }
func (r StampMicro) Format() string                       { return time.StampMicro }

// StampNano = Jan _2 15:04:05.000000000
type StampNano string

func StringStampNano(input time.Time) String[StampNano] { return String[StampNano](input) }
func (r StampNano) Format() string                      { return time.StampNano }

// DateTime = 2006-01-02 15:04:05
type DateTime string //nolint:revive

func StringDateTime(input time.Time) String[DateTime] { return String[DateTime](input) }
func (r DateTime) Format() string                     { return time.DateTime }

// DateOnly = 2006-01-02
type DateOnly string //nolint:revive

func StringDateOnly(input time.Time) String[DateOnly] { return String[DateOnly](input) }
func (r DateOnly) Format() string                     { return time.DateOnly }

// TimeOnly = 15:04:05
type TimeOnly string

func StringTimeOnly(input time.Time) String[TimeOnly] { return String[TimeOnly](input) }
func (r TimeOnly) Format() string                     { return time.TimeOnly }

type String[T StringFormat] time.Time

// Time returns the underlying time.Time value.
func (s String[T]) Time() time.Time { return time.Time(s) }

// String implements [fmt.Stringer].
func (s String[T]) String() string {
	var genericType T
	return s.Time().Format(genericType.Format())
}

// UnmarshalJSON implements [encoding/json.Unmarshaler].
func (s *String[T]) UnmarshalJSON(input []byte) error {
	var (
		genericType T
		cleaned     = strings.Trim(string(input), `"`)
	)

	parsed, err := time.Parse(genericType.Format(), cleaned)
	if err != nil {
		return fmt.Errorf("failed to parse time in %q format: %w", genericType.Format(), err)
	}

	*s = String[T](parsed)

	return nil
}

// MarshalJSON implements [encoding/json.Marshaler].
func (s String[T]) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

// EncodeValues implements [github.com/google/go-querystring/query.Encoder] (github.com/google/go-querystring/query)
func (s String[T]) EncodeValues(key string, values *url.Values) error {
	if values == nil {
		values = &url.Values{}
	} else if *values == nil {
		*values = url.Values{}
	}

	values.Add(key, s.String())

	return nil
}
