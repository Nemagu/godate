package godate

import (
	"database/sql/driver"
	"fmt"
	"time"
)

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

func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.String(), nil
}
