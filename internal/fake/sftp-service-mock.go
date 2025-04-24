package fake

import (
	"github.com/pkg/sftp"
	"github.com/stretchr/testify/mock"
)

type MockSFTPService struct {
	mock.Mock
}

func (mock *MockSFTPService) ConnectClient() (*sftp.Client, error) {
	result := mock.Called()
	return result.Get(0).(*sftp.Client), result.Error(1)
}
func (mock *MockSFTPService) CloseClient() {
	mock.Called()
}
