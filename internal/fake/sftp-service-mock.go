package fake

import (
	"poc-auto-read-ec/models"

	"github.com/pkg/sftp"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/ssh"
)

type MockSFTPService struct {
	mock.Mock
}

func (mock *MockSFTPService) ConnectClient(conn *ssh.Client) (*sftp.Client, error) {
	result := mock.Called()
	return result.Get(0).(*sftp.Client), result.Error(1)
}
func (mock *MockSFTPService) CloseClient() {
	mock.Called()
}

func (mock *MockSFTPService) GetAllCSVFile() ([]models.CSVRawFile, error) {
	result := mock.Called()
		return result.Get(0).([]models.CSVRawFile), nil
}

func (mock *MockSFTPService) getFileContent(filename string) ([]byte, error) {
	result := mock.Called()
	return result.Get(0).([]byte), nil
}

func (mock *MockSFTPService) ParseCSVToPerson(data []byte) ([]models.Person, error) {
	result := mock.Called()
	return result.Get(0).([]models.Person), nil

}
func (mock *MockSFTPService) TransformPersonToGorm(listPerson []models.Person) ([]models.GormPerson, error) {
	result := mock.Called()
	return result.Get(0).([]models.GormPerson), nil
}
