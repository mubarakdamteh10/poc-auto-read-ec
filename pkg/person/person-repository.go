package person

type IPersonRepository interface{}

type personRepository struct{}

func NewPersonRepository() IPersonRepository {
	return &personRepository{}
}
