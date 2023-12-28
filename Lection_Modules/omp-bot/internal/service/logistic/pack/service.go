package pack

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Subdomain {
	return allEntities
}

func (s *Service) Get(idx int) (*Subdomain, error) {
	if idx <= len(allEntities) && !(idx < 0) {
		return &allEntities[idx], nil
	}
	return nil, errors.New("index is wrong")
}
