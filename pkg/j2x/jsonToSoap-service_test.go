package j2x

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
