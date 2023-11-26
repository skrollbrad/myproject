package person

import (
	"example/models"
)

type storage struct {
	person map[string]models.Person
}

func NewStorage() *storage {

	m := make(map[string]models.Person)

	return &storage{m}

}
func (s *storage) AddUser(p models.Person) error {

	s.person[p.Id] = p

	return nil
}
func (s *storage) GetUser(id string) (models.Person, error) {

	return s.person[id], nil
}
