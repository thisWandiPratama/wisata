package tourist

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]TouristSite, error)
	Search(name string) ([]TouristSite, error)
	FindAllByCategory(category_id int) ([]TouristSite, error)
	Save(tourist TouristSite) (TouristSite, error)
	Update(tourist TouristSite) (TouristSite, error)
	Delete(id int) (TouristSite, error)
	FindByID(ID int) (TouristSite, error)
	FindAllGallery(touristid int) ([]Gallery, error)
	SaveGallery(gallery Gallery) (Gallery, error)
	DeleteGallery(id int) (Gallery, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]TouristSite, error) {
	var tourist []TouristSite

	err := r.db.Find(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) FindAllByCategory(category_id int) ([]TouristSite, error) {
	var tourist []TouristSite

	err := r.db.Where("category_id", category_id).Find(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) Save(tourist TouristSite) (TouristSite, error) {
	err := r.db.Create(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) Update(tourist TouristSite) (TouristSite, error) {
	err := r.db.Save(&tourist).Error

	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) Delete(id int) (TouristSite, error) {
	var tourist TouristSite
	err := r.db.Delete(&TouristSite{}, id).Error

	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) FindByID(ID int) (TouristSite, error) {
	var tourist TouristSite

	err := r.db.Where("id = ?", ID).Find(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

// gallery
func (r *repository) FindAllGallery(touristid int) ([]Gallery, error) {
	var tourist []Gallery

	err := r.db.Find(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) SaveGallery(tourist Gallery) (Gallery, error) {
	err := r.db.Create(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) DeleteGallery(id int) (Gallery, error) {
	var tourist Gallery
	fmt.Println(id)
	err := r.db.Where("id=?", id).Delete(&Gallery{}, id).Error

	if err != nil {
		return tourist, err
	}

	return tourist, nil
}

func (r *repository) Search(name string) ([]TouristSite, error) {
	var tourist []TouristSite
	dataName := "%" + name + "%"
	err := r.db.Where("name LIKE ?", dataName).Find(&tourist).Error
	if err != nil {
		return tourist, err
	}

	return tourist, nil
}
