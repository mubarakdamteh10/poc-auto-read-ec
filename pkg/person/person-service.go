package person

type IPersonService interface{}

type personService struct{}

func NewPersonService() IPersonService {
	return &personService{}
}
