package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(x); err != nil {
		return err
	}

	if decoder.More() {
		return errors.New("request body must only contain a single JSON object")
	}

	return nil
}
