package godate

import "encoding/json"

func (d Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Date) UnmarshalText(data []byte) error {
	temp, err := FromString(string(data))
	if err != nil {
		return err
	}
	*d = temp
	return nil
}

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
