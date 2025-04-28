package process

import (
	"poc-auto-read-ec/pkg/person"
	"poc-auto-read-ec/pkg/sftp"
)

type IProcessService interface {
	ProcessAutoReadEC() error

	ProcessConvertJsonToXML() error
}

type processService struct {
	sftpService   sftp.ISFTPService
	personService person.IPersonService
}

func NewProcessService() IProcessService {
	return &processService{
		sftpService:   sftp.NewSFTPService(),
		personService: person.NewPersonService(),
	}
}

func (service *processService) ProcessAutoReadEC() error {

	listRawCSV, err := service.sftpService.GetAllCSVFile()
	if err != nil {
		return err
	}

	listPerson, err := service.sftpService.ParseCSVToListRaw(listRawCSV)
	if err != nil {
		return err
	}

	listGormPerson, err := service.sftpService.TransformPersonToGorm(listPerson)
	if err != nil {
		return err
	}

	err = service.personService.SavePersonsToDB(listGormPerson)
	if err != nil {
		return err
	}

	return nil
}

func (service *processService) ProcessConvertJsonToXML() error {
	return nil
}
