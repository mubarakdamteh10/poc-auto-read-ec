package sftp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSFTPService(t *testing.T) {
	// os.Setenv("SFTP_HOST", "test")
	// os.Setenv("SFTP_PORT", "test")
	// os.Setenv("SFTP_USERNAME", "test")
	// os.Setenv("SFTP_PASSWORD", "test")
	// os.Setenv("SFTP_BASEPATH", "test")
	service := NewSFTPService()
	v := reflect.Indirect(reflect.ValueOf(service))
	for index := 0; index < v.NumField(); index++ {
		if v.Type().Field(index).Name != "client" {
			assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
		}
	}
}