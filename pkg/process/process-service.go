package process

import "poc-auto-read-ec/pkg/sftp"

type IProcessService interface {
	ProcessAutoReadEC() error

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
