package person

import "poc-auto-read-ec/models"

type IPersonService interface {
	SavePersonsToDB(list []models.GormPerson) error
}

type personService struct {
	repository IPersonRepository
}

func NewPersonService() IPersonService {
	return &personService{
		repository: NewPersonRepository(),
	}
}

func (service *personService) SavePersonsToDB(list []models.GormPerson) error {
	err := service.repository.InsertPersonToDB(list)
	if err != nil {
		return err
	}
	return nil
}
