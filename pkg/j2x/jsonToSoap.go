package j2x

type IJsonToSoap interface{}

type jsonToSoap struct{}

func NewJsonToSoap() IJsonToSoap {
	return &jsonToSoap{}
}
