package sftp

import (
	"github.com/pkg/sftp"
)

type ISFTPService interface {

	// CloseClient closes the SFTP client connection
	//	input:
	//	- none
	//	output:
	//	- none
	CloseClient()

	// ConnectClient establishes a connection to the SFTP server
	//	input:
	//	- none
	//	output:
	//	- *sftp.Client: a pointer to the SFTP client
	//	- error: an error if the connection fails
	ConnectClient() (*sftp.Client, error)
}

type sftpService struct {
	client *sftp.Client
}

func NewSFTPService() ISFTPService {
	return &sftpService{}
}

func (service *sftpService) ConnectClient() (*sftp.Client, error) {
	return nil, nil
}

func (service *sftpService) CloseClient() {
	service.client.Close()
}
