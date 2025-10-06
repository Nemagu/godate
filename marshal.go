package godate

import (
	_ "encoding"
	"encoding/json"
)

// MarshalText implements the [encoding.TextMarshaler] interface.
// The output matches that of calling the [Date.String] method.
func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

// MarshalJSON implements the [json.Marshaler] interface.
// The date must be in the ISO 8601 format.
func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
// The date must be in the ISO 8601 format.
func (d *Date) UnmarshalText(data []byte) error {
	temp, err := FromString(string(data))
	if err != nil {
		return err
	}
	*d = temp
	return nil
}

// UnmarshalJSON implements the [json.Unmarshaler] interface.
// The date must be in the ISO 8601 format.
func (d *Date) UnmarshalJSON(data []byte) error {
	temp := ""
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	date, err := FromString(temp)
	if err != nil {
		return err
	}
	*d = date
	return nil
}
