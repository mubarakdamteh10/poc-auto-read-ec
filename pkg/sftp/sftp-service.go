package sftp

import (
	"fmt"

	"github.com/pkg/sftp"
)

type ISFTPService interface {

	// CloseClient closes the SFTP client connection
	//	input:
	//	- none
	//	output:
	//	- none
	CloseClient()
}

type sftpService struct{}

func NewSFTPService() ISFTPService {
	return &sftpService{}
}

func (service *sftpService) ConnectClient() (*sftp.Client, error) {
	return nil, nil
}

func (service *sftpService) CloseClient() {
	fmt.Println("close")
}
