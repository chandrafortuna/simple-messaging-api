package message

//MessageRepository interface
type Repository interface {
	GetAll() ([]*Message, error)
	Save(m *Message) error
}

type TempRepository struct {
	messageCollection []*Message
}

func (r *TempRepository) GetAll() ([]*Message, error) {
	return r.messageCollection, nil
}

func (r *TempRepository) Save(m *Message) error {
	r.messageCollection = append(r.messageCollection, m)
	return nil
}

func NewRepository(m []*Message) (r Repository) {
	r = &TempRepository{m}
	return
}
