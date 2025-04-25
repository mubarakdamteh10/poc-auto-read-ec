package person

import (
	"errors"
	"poc-auto-read-ec/models"
)

type IPersonRepository interface {
	// InsertPersonToDB inserts list of person into the database
	// input :
	//	- list : []models.GormPerson
	// output :
	//	- error : error
	InsertPersonToDB(list []models.GormPerson) error
}

type personRepository struct{}

func NewPersonRepository() IPersonRepository {
	return &personRepository{}
}

func (repo *personRepository) InsertPersonToDB(list []models.GormPerson) error {
	return errors.New("waiting for implement")
}
