package j2x

import (
	"errors"
	"poc-auto-read-ec/models"
)

type IJsonToSoap interface {
	CreateSoapData() (*models.SoapData, error)
	ExtractJSON() error
}

type jsonToSoap struct{}

func (service *jsonToSoap) CreateSoapData() (*models.SoapData, error) {
	return &models.SoapData{}, nil

}

func NewJsonToSoap() IJsonToSoap {
	return &jsonToSoap{}
}

func (service *jsonToSoap) ExtractJSON() error {
	return errors.New("waiting for implement")
}
