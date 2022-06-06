package service

import (
	"fmt"
	"net/http"

	"BE/config"
	"BE/domain"
	"BE/helper"
	"BE/model"
)

type svcUser struct {
	c config.Config
	repo domain.AdapterRepositoryUser
}

func (s *svcUser) CreateUserService(user model.User) error {
	return s.repo.CreateUsers(user)
}

func (s *svcUser) UpdateUserService(id, idToken int, user model.User) error {
	fmt.Println(id, idToken)
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneByID(id, user)
}

func (s *svcUser) GetAllUsersService() []model.User {
	return s.repo.GetAll()
}

func (s *svcUser) GetUserByID(id int) (model.User, error) {
	return s.repo.GetOneByID(id)
}

func (s *svcUser) LoginUser(email, password string) (string, int) {
	user, _ := s.repo.GetOneByEmail(email)

	if (user.Password != password) || (user == model.User{}) {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(user.ID, user.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *svcUser) DeleteByID(id int) error {
	return s.repo.DeleteByID(id)
}

func NewServiceUser(repo domain.AdapterRepositoryUser, c config.Config) domain.AdapterService {
	return &svcUser{
		repo: repo,
		c: c,
	}
}