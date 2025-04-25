package sftp

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"poc-auto-read-ec/environment"
	"poc-auto-read-ec/models"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
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
	ConnectClient(*ssh.Client) (*sftp.Client, error)

	// GetAllCSVFile retrieves all CSVRawFile entries from the SFTP server
	//	input:
	//	- none
	//	output:
	//	- []models.CSVRawFile: a slice of CSVRawFile objects found
	//	- error: an error if the retrieval fails
	GetAllCSVFile() ([]models.CSVRawFile, error)

	// ParseCSVToPerson parses raw CSV data into a slice of Person structs
	//	input:
	//	- data []byte: raw CSV file content as byte slice, including header row
	//	output:
	//	- []models.Person: a slice of Person structs parsed from CSV
	//	- error: an error if parsing fails or headers/columns mismatch
	ParseCSVToPerson(data []byte) ([]models.Person, error)

	// TransformPersonToGorm transforms a slice of Person structs
	// input :
	//	- listPerson []models.Person
	// output:
	//	- []models.GormPerson
	//	- error
	TransformPersonToGorm(listPerson []models.Person) ([]models.GormPerson, error)
}

type ISftpClientFactory interface {
	NewClient(conn *ssh.Client) (*sftp.Client, error)
}

type sftpService struct {
	client *sftp.Client
	// clientFactory ISftpClientFactory
	// conn          *ssh.Client
}

func NewSFTPService() ISFTPService {
	return &sftpService{}
}

func NewClient(conn *ssh.Client) (*sftp.Client, error) {

	if conn == nil {
		return nil, errors.New("ssh client is nil")
	}

	return sftp.NewClient(conn)

}

// func (service *sftpService) ConnectClient() (*sftp.Client, error) {
// 	client, err := service.clientFactory.NewClient(service.conn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Close()
// 	return client, nil
// }

func (service *sftpService) ConnectClient(conn *ssh.Client) (*sftp.Client, error) {
	if conn == nil {
		return nil, errors.New("connection is not specified")
	}

	return sftp.NewClient(conn)
}

func (service *sftpService) CloseClient() {
	if service.client != nil {
		service.client.Close()
	}
}

func (service *sftpService) GetAllCSVFile() ([]models.CSVRawFile, error) {
	config := environment.GetSFTPConfiguration()
	directory := config.BasePath

	listFile, err := service.client.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	listCSVFile := []models.CSVRawFile{}

	for _, fileInfo := range listFile {
		if !fileInfo.IsDir() && strings.EqualFold(filepath.Ext(fileInfo.Name()), ".csv") {
			absFileName := fmt.Sprintf("%s/%s", directory, fileInfo.Name())

			content, err := service.getFileContent(absFileName)
			if err != nil {
				return nil, err
			}

			csvFile := models.CSVRawFile{
				FileName: fileInfo.Name(),
				RawFile:  content,
			}

			listCSVFile = append(listCSVFile, csvFile)
		}
	}

	return listCSVFile, nil
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

func (service *sftpService) ParseCSVToPerson(data []byte) ([]models.Person, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	for value := range header {
		header[value] = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(header[value]), " ", "_"))
	}

	var people []models.Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %w", err)
		}
		if len(record) != len(header) {
			return nil, fmt.Errorf("record length mismatch")
		}

		person := models.Person{}
		for index, value := range header {
			recordValue := strings.TrimSpace(record[index])
			switch value {
			case "first_name":
				person.FirstName = recordValue
			case "last_name":
				person.LastName = recordValue
			case "email":
				person.Email = recordValue
			case "phone_number":
				person.PhoneNumber = recordValue
			case "date_of_birth":
				person.DateOfBirth = recordValue
			case "address":
				person.Address = recordValue
			}
		}

		people = append(people, person)
	}

	return people, nil
}

func (service *sftpService) TransformPersonToGorm(listPerson []models.Person) ([]models.GormPerson, error) {
	listGorm := []models.GormPerson{}
	for _, person := range listPerson {
		gormPerson := models.GormPerson{
			FirstName:   person.FirstName,
			LastName:    person.LastName,
			Email:       person.Email,
			PhoneNumber: person.PhoneNumber,
			DateOfBirth: person.DateOfBirth,
			Address:     person.Address,
		}
		listGorm = append(listGorm, gormPerson)
	}
	return listGorm, nil
}
