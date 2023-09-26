package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func Call(url string) (*http.Response, error) {
	return http.Get(url)
}

func DecoderReader(in io.Reader, out interface{}) error {
	return json.NewDecoder(in).Decode(&out)
}
