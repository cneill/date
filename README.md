# date

[![Go Reference](https://pkg.go.dev/badge/github.com/cneill/date.svg)](https://pkg.go.dev/github.com/cneill/date)

Some date utilities for Go.

## Example

```golang
type Data struct {
    Timestamp date.String[date.RFC3339] `json:"time_stamp"`
}

func (d *Data) After(input time.Time) bool {
    return d.Timestamp.Time().After(input)
}
```
