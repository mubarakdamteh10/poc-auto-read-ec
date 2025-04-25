package person

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonRepository(t *testing.T) {
	repository := NewPersonRepository()

	v := reflect.Indirect(reflect.ValueOf(repository))
	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}

	assert.NotNil(t, repository, "Service should not be nil")
}
