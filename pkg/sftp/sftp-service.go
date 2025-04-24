package sftp

import (
	"fmt"
	"io"
	"poc-auto-read-ec/environment"
	"poc-auto-read-ec/models"

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

	// GetAllFileEC retrieves all XMLRawFile entries from the SFTP server
	//	input:
	//	- none
	//	output:
	//	- []models.XMLRawFile: a slice of XMLRawFile objects found
	//	- error: an error if the retrieval fails
	GetAllFileEC() ([]models.XMLRawFile,error)
}

type sftpService struct{
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

func (service *sftpService) GetAllFileEC() ([]models.XMLRawFile,error){

	config := environment.GetSFTPConfiguration()

	directory := config.BasePath
	listFile, err := service.client.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	listRawFile := []models.XMLRawFile{}

	for _, fileInfo := range listFile {
		if !fileInfo.IsDir() {
			absFileName := fmt.Sprintf("%s/%s", directory, fileInfo.Name())

			content, err := service.getFileContent(absFileName)
			if err != nil {
				return nil, err
			}

			rawFile := models.XMLRawFile{
				FileName:         fileInfo.Name(),
				RawFile:          content,
			}

			listRawFile = append(listRawFile, rawFile)
		}
	}

	return  listRawFile, nil
}

func (service *sftpService) getFileContent(filename string) ([]byte, error) {
	file, err := service.client.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}

	return data, nil
}