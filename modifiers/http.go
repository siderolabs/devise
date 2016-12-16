package modifiers

import (
	"io/ioutil"
	"net/http"
)

type HTTP struct{}

func NewHTTP() *HTTP {
	h := HTTP{}

	return &h
}

func (h *HTTP) Get(s string) (*string, error) {

	resp, err := http.Get(s)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s = string(body)

	return &s, nil
}
