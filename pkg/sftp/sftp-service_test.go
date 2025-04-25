package sftp

import (
	"os"
	"reflect"
	"testing"

	"github.com/pkg/sftp"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
)

type MockSFTPClientFactory struct {
	ReturnClient *sftp.Client
	ReturnError  error
}
type IMockSFTPClientInterface interface {
	Close() error
	ReadDir(p string) ([]os.FileInfo, error)
	Open(path string) (*sftp.File, error)
}


func (m *MockSFTPClient) NewClient(conn *ssh.Client) (*sftp.Client, error) {
	return nil, nil
}
func TestNewSFTPService(t *testing.T) {
	// os.Setenv("SFTP_HOST", "test")
	// os.Setenv("SFTP_PORT", "test")
	// os.Setenv("SFTP_USERNAME", "test")
	// os.Setenv("SFTP_PASSWORD", "test")
	// os.Setenv("SFTP_BASEPATH", "test")
	service := NewSFTPService()
	v := reflect.Indirect(reflect.ValueOf(service))
	for index := 0; index < v.NumField(); index++ {
		name := v.Type().Field(index).Name
		if name != "client" {
			assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
		}
	}
}

// func (m *MockSFTPClient) Close() error {
// 	return nil
// }

// func (m *MockSFTPClient) ReadDir(p string) ([]os.FileInfo, error) {
// 	return []os.FileInfo{}, nil
// }
// func (m *MockSFTPClient) Open(path string) (*sftp.File, error) {
// 	return nil, nil
// }

type MockSFTPClient struct {
	sftp.Client
}
func TestConnectClient_Failed(t *testing.T) {


	service := &sftpService{
 
	}

	_, err := service.ConnectClient(nil)

	if err == nil {
		t.Errorf("Expected error 'mock connection error', got: %v", err)
	}

}
