package tourist

import "fmt"

type Service interface {
	FindAll() ([]TouristSite, error)
	FindByID(id int) (TouristSite, error)
	Search(name string) ([]TouristSite, error)
	FindAllByCategory(category_id int) ([]TouristSite, error)
	Save(input CreateTouristInput) (TouristSite, error)
	Update(input UpdateTouristInput) (TouristSite, error)
	Delete(id int) (TouristSite, error)
	FindAllGallery(touristid int) ([]Gallery, error)
	SaveGallery(gallery Gallery) (Gallery, error)
	DeleteGallery(id int) (Gallery, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]TouristSite, error) {
	tourists, err := s.repository.FindAll()
	if err != nil {
		return tourists, err
	}

	return tourists, nil
}

func (s *service) FindByID(id int) (TouristSite, error) {
	tourists, err := s.repository.FindByID(id)
	if err != nil {
		return tourists, err
	}

	return tourists, nil
}

func (s *service) Search(name string) ([]TouristSite, error) {
	tourists, err := s.repository.Search(name)
	if err != nil {
		return tourists, err
	}

	return tourists, nil
}

func (s *service) FindAllByCategory(category_id int) ([]TouristSite, error) {
	tourists, err := s.repository.FindAllByCategory(category_id)
	if err != nil {
		return tourists, err
	}

	return tourists, nil
}

func (s *service) Save(input CreateTouristInput) (TouristSite, error) {
	tourist := TouristSite{}
	tourist.CategoryID = input.CategoryID
	tourist.Name = input.Name
	tourist.Address = input.Address
	tourist.Description = input.Description
	tourist.Email = input.Email
	tourist.Latitude = input.Latitude
	tourist.Longitude = input.Longitude
	tourist.Website = input.Website
	tourist.LinkVideo = input.LinkVideo
	tourist.Phone = input.Phone
	tourist.ImagePrimary = input.ImagePrimary

	newCategory, err := s.repository.Save(tourist)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) Update(input UpdateTouristInput) (TouristSite, error) {
	tourist, err := s.repository.FindByID(input.ID)
	fmt.Println(tourist)
	if err != nil {
		return tourist, err
	}

	tourist.CategoryID = input.CategoryID
	tourist.Name = input.Name
	tourist.Address = input.Address
	tourist.Description = input.Description
	tourist.Email = input.Email
	tourist.Latitude = input.Latitude
	tourist.Longitude = input.Longitude
	tourist.Website = input.Website
	tourist.LinkVideo = input.LinkVideo
	tourist.Phone = input.Phone
	tourist.ImagePrimary = input.ImagePrimary

	updatedTourist, err := s.repository.Update(tourist)
	if err != nil {
		return updatedTourist, err
	}

	return updatedTourist, nil
}

func (s *service) Delete(id int) (TouristSite, error) {
	deleteTourist, err := s.repository.Delete(id)
	if err != nil {
		return deleteTourist, err
	}

	return deleteTourist, nil
}

// gallery
func (s *service) FindAllGallery(touristid int) ([]Gallery, error) {
	tourists, err := s.repository.FindAllGallery(touristid)
	if err != nil {
		return tourists, err
	}

	return tourists, nil
}

func (s *service) SaveGallery(input Gallery) (Gallery, error) {
	tourist := Gallery{}
	tourist.TouristID = input.TouristID
	tourist.Avatar = input.Avatar

	newCategory, err := s.repository.SaveGallery(tourist)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) DeleteGallery(id int) (Gallery, error) {
	deleteTourist, err := s.repository.DeleteGallery(id)
	if err != nil {
		return deleteTourist, err
	}
	return deleteTourist, nil
}
