package process

import (
	"errors"
	"poc-auto-read-ec/internal/fake"
	"poc-auto-read-ec/models"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProcessService(t *testing.T) {
	service := NewProcessService()

	v := reflect.Indirect(reflect.ValueOf(service))

	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}
}

func TestProcessAutoReadEC_CallsProcessAutoReadEC_Success(t *testing.T) {
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

func TestProcessAutoReadEC_CallsProcessAutoReadEC_Failed(t *testing.T) {
	mockSftpService := &fake.MockSFTPService{}
	mockSftpService.Mock.On("GetAllCSVFile").Return(errors.New("cannot read the csv files from the server"))
	service := &processService{
		sftpService: mockSftpService,
	}
	err := service.ProcessConvertJsonToXML()
	

	if(err == nil){
		t.Logf("Expected error, got %v",err)
	}
}
