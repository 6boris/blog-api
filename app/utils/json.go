package utils

import (
	"encoding/json"
	"log"
)

func JsonMarshalIndent(v interface{}, prefix string, indent string) string {
	data, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		log.Println(err.Error())
	}
	return string(data)
}
