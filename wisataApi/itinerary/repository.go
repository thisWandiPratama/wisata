package itinerary

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Itinerary, error)
	FindAllByIDUser(id int) ([]Itinerary, error)
	Save(itinerary Itinerary) (Itinerary, error)
	Update(itinerary Itinerary) (Itinerary, error)
	Delete(id int) (Itinerary, error)
	FindByID(ID int) (Itinerary, error)

	// timeline
	SaveTimeline(timeline Timeline) (Timeline, error)
	UpdateTimeline(timeline Timeline) (Timeline, error)
	DeleteTimeline(id int) (Timeline, error)
	FindByIDTimeline(ID int) (Timeline, error)
	FindAllByIDTimeline(ID int) ([]Timeline, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Itinerary, error) {
	var itinerary []Itinerary

	err := r.db.Find(&itinerary).Error
	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) FindAllByIDUser(id int) ([]Itinerary, error) {
	var itinerary []Itinerary

	err := r.db.Where("user_id=?", id).Find(&itinerary).Error
	if err != nil {
		return itinerary, err
	}
	return itinerary, nil
}

func (r *repository) SaveTimeline(itinerary Timeline) (Timeline, error) {
	err := r.db.Create(&itinerary).Error
	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) UpdateTimeline(itinerary Timeline) (Timeline, error) {
	err := r.db.Save(&itinerary).Error

	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) DeleteTimeline(id int) (Timeline, error) {
	var itinerary Timeline
	err := r.db.Delete(&Itinerary{}, id).Error

	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) FindByIDTimeline(ID int) (Timeline, error) {
	var itinerary Timeline

	err := r.db.Where("itinerary_id = ?", ID).Find(&itinerary).Error
	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) FindAllByIDTimeline(ID int) ([]Timeline, error) {
	var itinerary []Timeline

	err := r.db.Where("itinerary_id = ?", ID).Find(&itinerary).Error
	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

// time line

func (r *repository) Save(itinerary Itinerary) (Itinerary, error) {
	err := r.db.Create(&itinerary).Error
	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) Update(itinerary Itinerary) (Itinerary, error) {
	err := r.db.Save(&itinerary).Error

	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) Delete(id int) (Itinerary, error) {
	var itinerary Itinerary
	err := r.db.Delete(&Itinerary{}, id).Error

	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}

func (r *repository) FindByID(ID int) (Itinerary, error) {
	var itinerary Itinerary

	err := r.db.Where("id = ?", ID).Find(&itinerary).Error
	if err != nil {
		return itinerary, err
	}

	return itinerary, nil
}
