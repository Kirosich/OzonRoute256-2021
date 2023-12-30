package pack

import (
	"errors"

	logistic "github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/model/logistic"
)

var allEntities = []logistic.Pack{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
	{Title: "six"},
	{Title: "seven"},
	{Title: "eight"},
	{Title: "nine"},
	{Title: "ten"},
	{Title: "eleven"},
	{Title: "twelve"},
	{Title: "thirteen"},
	{Title: "fourteen"},
	{Title: "fifteen"},
	{Title: "sixteen"},
	{Title: "seventeen"},
}

type PackService interface {
	Describe(packID uint64) (*logistic.Pack, error)
	List(cursor uint64, limit uint64) ([]logistic.Pack, error)
	Create(pack logistic.Pack) (uint64, error)
	Update(packID uint64, pack logistic.Pack) error
	Remove(packID uint64) (bool, error)
}

type DummyPackService struct{}

func NewDummyPackService() *DummyPackService {
	return &DummyPackService{}
}

func (s *DummyPackService) Describe(packID uint64) (*logistic.Pack, error) {

	if int(packID) <= len(allEntities) && !(int(packID) < 0) {
		getElem := allEntities[packID]
		return &getElem, nil
	}
	return nil, errors.New("index is wrong")
}

func (s *DummyPackService) List(cursor uint64, limit uint64) ([]logistic.Pack, error) {

	var allEntitiesSlice []logistic.Pack
	rlen := int(cursor + limit) // Последний элемент, который будет брать функция

	for i := int(cursor); i < rlen; i++ {
		if i > len(allEntities)-1 {
			break
		}

		allEntitiesSlice = append(allEntitiesSlice, allEntities[i])

	}
	return allEntitiesSlice, nil
}

func (s *DummyPackService) Create(pack logistic.Pack) (uint64, error) {
	allEntities = append(allEntities, pack)
	return uint64(len(allEntities)), nil
}

func (s *DummyPackService) Update(packID uint64, pack logistic.Pack) error {
	if packID < uint64(len(allEntities)) {
		allEntities[packID] = pack
		return nil
	}

	return errors.New("packid more than len of list")
}

func (s *DummyPackService) Remove(packID uint64) (bool, error) {
	if int(packID) < len(allEntities) {
		var newAllEntities []logistic.Pack
		newAllEntities = append(newAllEntities, allEntities[0:int(packID)]...)
		newAllEntities = append(newAllEntities, allEntities[int(packID)+1:]...)
		allEntities = newAllEntities

		return true, nil
	}

	return false, errors.New("wrong pack ID")
}
