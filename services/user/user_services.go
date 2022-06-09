package service

import (
	config "BE/configs"
	"BE/domain"
	"BE/helper/middleware"
	"BE/model"
	"fmt"
	"log"
	"net/http"
)

type svcUser struct {
	c    config.Config
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
	user, err := s.repo.Login(email, password)

	if err != nil {
		log.Println("Your Password Error", err)
		return "Your Password Error", http.StatusUnauthorized
	}

	token, _ := middleware.CreateToken(int(user.ID), int(user.RoleId), user.Username, s.c.JWT_KEY)
	return token, http.StatusOK
}

func (s *svcUser) DeleteByID(id int) error {
	return s.repo.DeleteByID(id)
}

func NewServiceUser(repo domain.AdapterRepositoryUser, c config.Config) domain.AdapterServiceUser {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
