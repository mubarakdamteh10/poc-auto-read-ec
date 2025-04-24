package environment

import (
	"os"
	"poc-auto-read-ec/models"
)

func GetSFTPConfiguration() models.SFTPConfiguration {
	host := os.Getenv("sftpHost")
	port := os.Getenv("sftpPort")
	username := os.Getenv("sftpUser")
	password := os.Getenv("sftpPassword")
	basePath := os.Getenv("sftpBasePath")

	return models.SFTPConfiguration{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		BasePath: basePath,
	}
}
