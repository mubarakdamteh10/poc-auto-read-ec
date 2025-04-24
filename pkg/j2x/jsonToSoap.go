package j2x

import "errors"

type IJsonToSoap interface {
	ExtractJSON() error
}

type jsonToSoap struct{}

func NewJsonToSoap() IJsonToSoap {
	return &jsonToSoap{}
}

func (service *jsonToSoap) ExtractJSON() error {
	return errors.New("waiting for implement")
}
