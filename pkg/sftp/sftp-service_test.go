package sftp

import (
	"os"
	"poc-auto-read-ec/models"
	"reflect"
	"testing"

	"github.com/pkg/sftp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	service := &sftpService{}
	_, err := service.ConnectClient(nil)

	if err == nil {
		t.Errorf("Expected error 'mock connection error', got: %v", err)
	}

}

func TestTransformPersonToGorm(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// Arrange
		mockListPerson := []models.Person{
			{
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
			{
				FirstName:   "Jane",
				LastName:    "MaiDee",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
		}

		service := NewSFTPService()

		// Act
		result, err := service.TransformPersonToGorm(mockListPerson)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, len(result), len(mockListPerson))
		expected := []models.GormPerson{
			{
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
			{
				FirstName:   "Jane",
				LastName:    "MaiDee",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
		}
		assert.Equal(t, expected, result)
	})

}

func TestParseCSVToPerson_Valid(t *testing.T) {
	t.Run("valid record", func(t *testing.T) {
	//Arrange
    input := `First Name,Last Name,Email,Phone Number,Date of Birth,Address
	John,Doe,john@example.com,1234567890,1990-01-01,123 Street`

	service :=  &sftpService{}

	//Act
    actualResult, err := service.ParseCSVToPerson([]byte(input))

	//Assert
    require.NoError(t, err)

	expectedResult := []models.Person{
        {
            FirstName:   "John",
            LastName:    "Doe",
            Email:       "john@example.com",
            PhoneNumber: "1234567890",
            DateOfBirth: "1990-01-01",
            Address:     "123 Street",
        },
	}
    assert.Equal(t, expectedResult, actualResult)
	})

	t.Run("failed to read record", func(t *testing.T) {
		//Arrange
        input := `First Name,Last Name,Email,Phone Number,Date of Birth,Address
		"John,Doe,john@example.com,1234567890,1990-01-01,123 Street`

		service :=  &sftpService{}

		//Act
        _, err := service.ParseCSVToPerson([]byte(input))

		//Assert
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "failed to read record")

    })
}


// func TestConnectClient_Success(t *testing.T) {
// 	service := &sftpService{}

// 	dummy := &ssh.Client{}
// 	func (d *ssh.Client) NewSession() (*Session, error) {
// 		return nil,nil
// 	}
// 	client, err := service.ConnectClient(dummy)

// 	if err != nil {
// 		t.Errorf("Expected error to nil, got: %v", err)
// 	}

// 	if client == nil {
// 		t.Errorf("Expected a client instance, got: %v", client)
// 	}

// }
