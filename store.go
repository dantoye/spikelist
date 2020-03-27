package main

type Store interface {
	FindMaker(nickname string) (string, error)
	AddMaker(nickname, code string) error
	UploadLevel(level *Level) error
	RateLevel(level *Level, maker string, rating float64) error
	GetLevel(code string) (*Level, error)
}

func NewStore() *store {
	return new(store)
}

type store struct{}

func (s *store) FindMaker(nickname string) (string, error) {
	return "XXX-XXX-XXX", nil
}
func (s *store) AddMaker(nickname, code string) error {
	return nil
}
func (s *store) UploadLevel(level *Level) error {
	return nil
}
func (s *store) RateLevel(level *Level, maker string, rating float64) error {
	return nil
}
func (s *store) GetLevel(code string) (*Level, error) {
	return &Level{Code: code}, nil
}
