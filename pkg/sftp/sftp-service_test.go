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
	service := NewSFTPService()
	v := reflect.Indirect(reflect.ValueOf(service))
	for index := 0; index < v.NumField(); index++ {
		name := v.Type().Field(index).Name
		if name != "client" {
			assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
		}
	}
}

type MockSFTPClient struct {
	sftp.Client
}

func TestTransformPersonToGorm(t *testing.T) {

	t.Run("success", func(t *testing.T) {
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

		result, err := service.TransformPersonToGorm(mockListPerson)

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

func TestExtractRawCSVToPerson(t *testing.T) {
	t.Run("valid record", func(t *testing.T) {
		input := `First Name,Last Name,Email,Phone Number,Date of Birth,Address
	John,Doe,john@example.com,1234567890,1990-01-01,123 Street`

		service := &sftpService{}

		actualResult, err := service.ExtractRawCSVToPerson([]byte(input))

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
		input := `First Name,Last Name,Email,Phone Number,Date of Birth,Address
		"John,Doe,john@example.com,1234567890,1990-01-01,123 Street`

		service := &sftpService{}

		_, err := service.ExtractRawCSVToPerson([]byte(input))

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to read record")

	})
}

func TestMapRecordToPerson(t *testing.T) {
	header := []string{
		"first_name", "last_name", "email", "phone_number", "date_of_birth", "address",
	}
	record := []string{
		"Jeng", "TestDee", "jengTestDee@example.com", "0981234567", "1988-11-06", "512 Bang Son Bangkok",
	}

	person, err := MapRecordToPerson(header, record)

	expected := models.Person{
		FirstName:   "Jeng",
		LastName:    "TestDee",
		Email:       "jengTestDee@example.com",
		PhoneNumber: "0981234567",
		DateOfBirth: "1988-11-06",
		Address:     "512 Bang Son Bangkok",
	}

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if person != expected {
		t.Errorf("expected %+v, got %+v", expected, person)
	}
}
