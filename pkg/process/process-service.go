package process

import "poc-auto-read-ec/pkg/sftp"

type IProcessService interface {
	// AutoReadEC reads the EC data from SFTP
	//	input:
	//	- none
	//	output:
	//	- error
	ProcessAutoReadEC() error

	// ProcessConvertJsonToXML converts JSON data to XML
	//	input:
	//	- none
	//	output:
	//	- error
	ProcessConvertJsonToXML() error
}

type processService struct {
	sftpService sftp.ISFTPService
}

func NewProcessService() IProcessService {
	return &processService{
		sftpService: sftp.NewSFTPService(),
	}
}

func (service *processService) ProcessAutoReadEC() error {

	_, err := service.sftpService.GetAllCSVFile()
	if err != nil {
		return err
	}

	// Calling function ConvertJsonToGORM

	// Calling function SaveTODB
	return nil
}

func (service *processService) ProcessConvertJsonToXML() error {
	return nil
}
