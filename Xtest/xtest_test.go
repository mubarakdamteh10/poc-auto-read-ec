package xtest

import (
	"os"
	"poc-auto-read-ec/models"
	"poc-auto-read-ec/pkg/person"
	"testing"
)

func XTestInsertPersonToDB(t *testing.T) {

	os.Setenv("DB_USER", "sa")
	os.Setenv("DB_PASSWORD", "Maibok!In0ng")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "1433")
	os.Setenv("DB_NAME", "master")
	repo := person.NewPersonRepository()

	testData := []models.GormPerson{
		{
			FirstName:   "Mighty",
			LastName:    "Sanchez",
			Email:       "Santomaru@fakeMail.com",
			PhoneNumber: "+66999999999",
			DateOfBirth: "1990-01-01",
			Address:     "123 Main St, City, Country",
		},
		{
			FirstName:   "PAeng",
			LastName:    "INong",
			Email:       "Maibok@fakeMail.com",
			PhoneNumber: "+66999999999",
			DateOfBirth: "1990-01-01",
			Address:     "123 Main St, City, Country",
		},
	}

	err := repo.InsertPersonToDB(testData)
	if err != nil {
		t.Errorf("Expected no error, got 4444 %v", err)
	}

}
