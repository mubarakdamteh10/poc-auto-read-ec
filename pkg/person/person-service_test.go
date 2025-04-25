package person

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonService(t *testing.T) {
	service := NewPersonService()

	v := reflect.Indirect(reflect.ValueOf(service))
	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}

	assert.NotNil(t, service, "Service should not be nil")
}
