package message

import "errors"

type Service struct {
	repo Repository
}

var ErrMessageSaveFailed = errors.New("Save Message Failed")
var ErrGetMessageFailed = errors.New("Get Message Failed")

func NewService(r Repository) Service {
	return Service{
		repo: r,
	}
}
func (s *Service) SetRepository(r Repository) {
	s.repo = r
}

func (s *Service) Send(msg string) (*Message, error) {
	m := &Message{
		Text: msg,
	}
	err := s.repo.Save(m)
	if err != nil {
		return nil, ErrMessageSaveFailed
	}

	return m, nil
}

func (s *Service) GetAll() ([]*Message, error) {
	msgCollection, err := s.repo.GetAll()
	if err != nil {
		return []*Message{}, ErrGetMessageFailed
	}
	return msgCollection, nil
}
