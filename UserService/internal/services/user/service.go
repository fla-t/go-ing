package user

import (
	"github.com/fla-t/go-ing/UserService/internal/domain/user"
)

type Service struct {
	repo user.UserRepository
}

func NewService(repo user.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) (id string) (*user.User, error) {

}
