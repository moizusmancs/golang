package auth

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(username, email, password string) (string, error) {

	_, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("email already exists")
	}

	hash, err := HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &User{
		Username: username,
		Email: email,
		PasswordHash: hash,
	}

	if err:=s.repo.CreateUser(user); err != nil {
		return "", err
	}

	return GenerateToken(user.ID)

}

func (s *Service) Login (email, password string) (string, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil{
		return "", err
	}

	if err := ComparePassword(user.PasswordHash, password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return GenerateToken(user.ID)
}