package sftp

import "fmt"

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

func (service *sftpService) CloseClient() {
	fmt.Println("close")
}
