package users

type UserRepository interface {
	GetById(id int) (*User, error)
	Save(user *User) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

func (s *UserService) GetById(id int) (*User, error) {
	return s.userRepository.GetById(id)
}

func (s *UserService) Save(user *User) error {
	return s.userRepository.Save(user)
}
