package demo

type Service struct {
	repository *Repository
}

func NewService() *Service {
	return &Service{repository: NewRepository()}
}

func (s *Service) GetUsers() []UserModel {
	ids := s.repository.GetUserIDs()
	return s.repository.GetUserByIDs(ids)
}
