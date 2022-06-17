package common

import "encoding/json"

func Swap(input, output interface{}) error {
	b, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, output)
}
