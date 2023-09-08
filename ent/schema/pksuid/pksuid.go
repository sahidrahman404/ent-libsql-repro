package pksuid

import (
	"database/sql/driver"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/segmentio/ksuid"
)

// ID implements a PKSUID - a prefixed ULID.
type ID string

// newKSUID returns a new KSUID for time.Now() using the default entropy source.
func newKSUID() string {
	k, err := ksuid.NewRandomWithTime(time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return k.String()
}

// MustNew returns a new PKSUID for time.Now() given a prefix. This uses the default entropy source.
func MustNew(prefix string) ID { return ID(prefix + newKSUID()) }

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return fmt.Errorf("pksuid: expected a value")
	}
	switch src := src.(type) {
	case string:
		*u = ID(src)
	case ID:
		*u = src
	default:
		return fmt.Errorf("pksuid: unexpected type, %T", src)
	}
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}
