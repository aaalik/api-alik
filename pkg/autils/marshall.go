package autils

import (
	"encoding/json"
)

func Marshall(obj interface{}) (body string) {
	bytes, err := json.Marshal(obj)
	if err == nil {
		body = string(bytes)
	}
	return
}
