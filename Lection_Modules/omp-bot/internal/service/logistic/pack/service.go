package pack

import (
	"errors"

	logistic "github.com/OzonRoute256-2021/Lection_Modules/omp-bot/internal/model/logistic"
)

var allEntitiesMap = map[int]logistic.Pack{
	0:  {Title: "one"},
	1:  {Title: "two"},
	2:  {Title: "three"},
	3:  {Title: "four"},
	4:  {Title: "five"},
	5:  {Title: "six"},
	6:  {Title: "seven"},
	7:  {Title: "eight"},
	8:  {Title: "nine"},
	9:  {Title: "ten"},
	10: {Title: "eleven"},
	11: {Title: "twelve"},
	12: {Title: "thirteen"},
	13: {Title: "fourteen"},
	14: {Title: "fifteen"},
	15: {Title: "sixteen"},
	16: {Title: "seventeen"}}

type PackService interface {
	Describe(packID uint64) (*logistic.Pack, error)
	List(cursor uint64, limit uint64) ([]logistic.Pack, error)
	Create(logistic.Pack) (uint64, error)
	Update(packID uint64, pack logistic.Pack) error
	Remove(packID uint64) (bool, error)
}

type DummyPackService struct{}

func NewService() *DummyPackService {
	return &DummyPackService{}
}

func (s *DummyPackService) List(cursor uint64, limit uint64) []logistic.Pack {

	var allEntitiesSlice []logistic.Pack
	rlen := int(cursor + limit) // Последний элемент, который будет брать функция

	for i := int(cursor); i < rlen; i++ {
		if int(cursor) > len(allEntitiesMap) {
			break
		}
		allEntitiesSlice = append(allEntitiesSlice, allEntitiesMap[i])
	}

	return allEntitiesSlice
}

func (s *DummyPackService) Describe(idx int) (*logistic.Pack, error) {
	if idx <= len(allEntitiesMap) && !(idx < 0) {
		getElem := allEntitiesMap[idx]
		return &getElem, nil
	}
	return nil, errors.New("index is wrong")
}
