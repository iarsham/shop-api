package common

import "encoding/json"

func TypeConverter[T any](data any) (*T, error) {
	var returnedType T
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataJson, &returnedType)
	if err != nil {
		return nil, err
	}
	return &returnedType, nil
}
