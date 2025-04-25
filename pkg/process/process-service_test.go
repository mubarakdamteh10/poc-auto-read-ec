package process

import (
	"github.com/stretchr/testify/assert"
	"poc-auto-read-ec/internal/fake"
	"poc-auto-read-ec/models"
	"reflect"
	"testing"
)

func TestNewProcessService(t *testing.T) {
	service := NewProcessService()

	v := reflect.Indirect(reflect.ValueOf(service))

	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}
}

func TestCallingConnectClient(t *testing.T) {
	mockSFTPService := &fake.MockSFTPService{}
	testFiles := []models.CSVRawFile{
		{FileName: "test-list.csv", RawFile: []byte{72, 101, 108, 108, 111}},
		{FileName: "test-list-2.csv", RawFile: []byte{72, 101, 108, 108, 108}},
	}
	mockSFTPService.Mock.On("GetAllCSVFile").Return(testFiles)
	service := &processService{
		sftpService: mockSFTPService,
	}

	service.ProcessAutoReadEC()
}
