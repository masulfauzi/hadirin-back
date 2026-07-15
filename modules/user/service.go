package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.FindAll()
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	return s.repo.FindByEmail(email)
}

func (s *Service) CreateUser(u *User) error {
	return s.repo.Create(u)
}
