package sftp

import "fmt"

type ISFTPService interface{}

type sftpService struct{}

func NewSFTPService() ISFTPService{
	return &sftpService{}
}

func (service *sftpService) CloseClient(){
	fmt.Println("close")
}