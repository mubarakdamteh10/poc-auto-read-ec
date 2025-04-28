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
	CloseClient()

	ConnectClient() (*sftp.Client, error)

	GetAllCSVFile() ([]models.CSVRawFile, error)

	GetFileContent(filename string) ([]byte, error)

	ExtractRawCSVToPerson(data []byte) ([]models.Person, error)

	ParseCSVToListRaw(files []models.CSVRawFile) ([]models.Person, error)

	TransformPersonToGorm(listPerson []models.Person) ([]models.GormPerson, error)
}

type ISftpClientFactory interface {
	NewClient(conn *ssh.Client) (*sftp.Client, error)
}

type sftpService struct {
	client *sftp.Client
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

func MapRecordToPerson(header, record []string) (models.Person, error) {

	person := models.Person{}

	for index, key := range header {
		recordValue := strings.TrimSpace(record[index])
		switch key {
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

	return person, nil
}

func (service *sftpService) ConnectClient() (*sftp.Client, error) {
	if service.client == nil {
		env := environment.GetSFTPConfiguration()

		config := &ssh.ClientConfig{
			User: env.Username,
			Auth: []ssh.AuthMethod{
				ssh.Password(env.Password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		ipPort := fmt.Sprintf("%s:%s", env.Host, env.Port)

		conn, err := ssh.Dial("tcp", ipPort, config)
		if err != nil {
			fmt.Printf("Failed to dial: %+v", err)
			return nil, err
		}

		client, err := sftp.NewClient(conn)
		if err != nil {
			fmt.Printf("Failed to create SFTP client: %+v", err)
			return nil, err
		}
		service.client = client
	}
	return service.client, nil
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

			content, err := service.GetFileContent(absFileName)
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

func (service *sftpService) ParseCSVToListRaw(files []models.CSVRawFile) ([]models.Person, error) {
	var people []models.Person

	for _, file := range files {
		persons, err := service.ExtractRawCSVToPerson(file.RawFile)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file %s: %w", file.FileName, err)
		}
		people = append(people, persons...)
	}

	return people, nil
}

func (service *sftpService) ExtractRawCSVToPerson(data []byte) ([]models.Person, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.TrimLeadingSpace = true

	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	for i := range header {
		header[i] = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(header[i]), " ", "_"))
	}

	people := []models.Person{}

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

		person, err := MapRecordToPerson(header, record)
		if err != nil {
			return nil, err
		}

		people = append(people, person)
	}

	return people, nil
}

func (service *sftpService) GetFileContent(filename string) ([]byte, error) {
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
