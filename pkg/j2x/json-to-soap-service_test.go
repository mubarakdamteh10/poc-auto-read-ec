package j2x

import (
	"github.com/stretchr/testify/assert"
	"poc-auto-read-ec/internal/fake"
	"reflect"
	"testing"
)

func TestNewJsonToSoap(t *testing.T) {
	service := NewJsonToSoap()

	v := reflect.Indirect(reflect.ValueOf(service))
	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}

	assert.NotNil(t, service, "Service should not be nil")
}

func TestExtractJSON(t *testing.T) {

	t.Run("success", func(t *testing.T) {
	})
}

func TestCreateSoapData(t *testing.T) {
	service := &fake.MockJSONToSoapService{}
	soapData, err := service.CreateSoapData()
	if soapData == nil {
		t.Logf("Expect an instance of soap data struct, got %v", soapData)

	}
	if err != nil {
		t.Logf("Expect no error, got %v", err)
	}
}
