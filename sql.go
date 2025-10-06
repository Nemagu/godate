package godate

import (
	_ "database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

// Scan implements the [sql.Scanner].
//
// Scan parse some value. If the value is nil the date will be zero.
// If the value holds string or []byte it calls [FromString] method;
// if the value holds [time.Time] it calls [FromTime] method;
// else Scan returns error.
func (d *Date) Scan(src any) error {
	var err error
	switch src := src.(type) {
	case nil:
		*d = 0
	case []byte:
		*d, err = FromString(string(src))
	case string:
		*d, err = FromString(src)
	case time.Time:
		*d = FromTime(src)
	default:
		err = fmt.Errorf("unsupported type: %T", src)
	}
	return err
}

// Value implements the [driver.Valuer].
//
// If [Date.IsZero] is true Value return nil as value for database;
// else it is false Value returns [Date.String] method's result.
func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.String(), nil
}
