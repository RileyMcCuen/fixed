package util

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	Error struct {
		Error string `json:"error"`
	}
)

func MarshalError(err error) string {
	log.Println("Marshalling error:", err.Error())
	e := Error{Error: err.Error()}
	data, _ := json.Marshal(e)
	return string(data)
}

func ErrorWrapper(desc string, err *error) {
	if err != nil && *err != nil {
		*err = fmt.Errorf("%w; %s", *err, desc)
	}
}
