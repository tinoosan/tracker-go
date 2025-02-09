package valueobjects

import (
	"time"
)

type DateTime struct {
	Value time.Time
}

func NewDateTime(value time.Time) *DateTime {
	return &DateTime{Value: value}
}

func (d *DateTime) String() string {
  return d.Value.Format("2006-01-02 15:04:05")
}

func (d *DateTime) DateString() string {
  return d.Value.Format("2006-01-02")
}

func (d *DateTime) TimeString() string {
  return d.Value.Format("15:04:05")
}

func (d *DateTime) Timestamp() int64 {
  return d.Value.Unix()
}

func (d *DateTime) Before(other *DateTime) bool {
  return d.Value.Before(other.Value)
}

func (d *DateTime) After(other *DateTime) bool {
  return d.Value.After(other.Value)
}
