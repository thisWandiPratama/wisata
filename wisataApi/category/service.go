package category

import "fmt"

type Service interface {
	FindAll() ([]Category, error)
	FindByID(id int) (Category, error)
	Save(input CreateCategoryInput) (Category, error)
	Update(input UpdateCategoryInput) (Category, error)
	Delete(id int) (Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Category, error) {
	categorys, err := s.repository.FindAll()
	if err != nil {
		return categorys, err
	}

	return categorys, nil
}

func (s *service) Save(input CreateCategoryInput) (Category, error) {
	campaign := Category{}
	campaign.Name = input.Name
	campaign.Avatar = input.Avatar

	newCategory, err := s.repository.Save(campaign)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) Update(input UpdateCategoryInput) (Category, error) {
	fmt.Println(input)
	category, err := s.repository.FindByID(input.ID)
	fmt.Println(category)
	if err != nil {
		return category, err
	}
	category.Name = input.Name
	category.Avatar = input.Avatar

	updatedCategory, err := s.repository.Update(category)
	if err != nil {
		return updatedCategory, err
	}

	return updatedCategory, nil
}

func (s *service) Delete(id int) (Category, error) {
	deleteCategory, err := s.repository.Delete(id)
	if err != nil {
		return deleteCategory, err
	}

	return deleteCategory, nil
}

func (s *service) FindByID(id int) (Category, error) {
	findCategory, err := s.repository.FindByID(id)
	if err != nil {
		return findCategory, err
	}

	return findCategory, nil
}
