package maybe

import (
	"bytes"
	"encoding/json"
)

var jsonNull = []byte("null")

func (o Maybe[T]) MarshalJSON() ([]byte, error) {
	if v, ok := o.Get(); ok {
		return json.Marshal(v)
	}
	
	return jsonNull, nil
}

func (o *Maybe[T]) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, jsonNull) {
		*o = None[T]()
		return nil
	}
	
	var v T
	
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	
	*o = Some(v)
	
	return nil
}
