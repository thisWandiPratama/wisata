package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	RegisterAdmin(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
	Delete(ID int) (User, error)
	GetAllUsers() ([]User, error)
	GetAllUsersByadmin() ([]User, error)
	UpdateUser(input FormUpdateUserInput) (User, error)
	Profile(input Profile) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	fmt.Println(input)
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.Avatar = "images/avatar_default.png"
	user.Role = "user"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	user1, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user1.Email == input.Email {
		fmt.Println("test")
		return user1, errors.New("email already exists")
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *service) RegisterAdmin(input RegisterUserInput) (User, error) {
	fmt.Println(input)
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.Avatar = "images/avatar_default.png"
	user.Role = "admin"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	user1, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user1.Email == input.Email {
		fmt.Println("test")
		return user1, errors.New("email already exists")
	}

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil

}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.Delete(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) GetAllUsersByadmin() ([]User, error) {
	users, err := s.repository.FindAllAdmin()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) Profile(input Profile) (User, error) {
	users, err := s.repository.FindByID(input.ID)
	if err != nil {
		return users, err
	}

	return users, nil
}
