package process

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

type processService struct{}

func NewProcessService() IProcessService {
	return &processService{}
}

func (service *processService) ProcessAutoReadEC() error {
	return nil
}

func (service *processService) ProcessConvertJsonToXML() error {
	return nil
}
