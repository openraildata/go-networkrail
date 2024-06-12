package smart

import (
	"encoding/json"
	"errors"
	"time"
)

type BerthOffset struct {
	time.Duration
}

func (b *BerthOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BerthOffset) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		b.Duration = time.Duration(value)
		return nil
	case string:
		// berth offsets are in seconds, suffix an 's' if it's not present
		if value[:len(value)-1] != "s" {
			value += "s"
		}
		var err error
		b.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}
