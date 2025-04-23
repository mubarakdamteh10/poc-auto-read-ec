package process

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProcessService(t *testing.T) {
	service := NewProcessService()

	v := reflect.Indirect(reflect.ValueOf(service))

	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}
}
