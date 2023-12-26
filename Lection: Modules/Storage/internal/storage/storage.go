package storage

type Storage struct {
	data string
}

func NewStorage() *Storage {
	newstorage := Storage{}
	return &newstorage
}
