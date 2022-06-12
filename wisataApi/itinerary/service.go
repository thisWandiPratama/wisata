package itinerary

type Service interface {
	FindAll() ([]Itinerary, error)
	FindAllByIDUser(id int) ([]Itinerary, error)
	Save(input CreateItineraryInput) (Itinerary, error)
	Update(input UpdateItineraryInput) (Itinerary, error)
	Delete(id int) (Itinerary, error)

	// timeline
	SaveTimeline(input CreateTimelineInput) (Timeline, error)
	UpdateTimeline(input UpdateTimelineInput) (Timeline, error)
	DeleteTimeline(id int) (Timeline, error)
	FindByIDTimeline(id int) ([]Timeline, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Itinerary, error) {
	itinerrays, err := s.repository.FindAll()
	if err != nil {
		return itinerrays, err
	}

	return itinerrays, nil
}

func (s *service) FindAllByIDUser(id int) ([]Itinerary, error) {
	itinerrays, err := s.repository.FindAllByIDUser(id)
	if err != nil {
		return itinerrays, err
	}

	return itinerrays, nil
}

func (s *service) Save(input CreateItineraryInput) (Itinerary, error) {
	itinerary := Itinerary{}
	itinerary.InitialLocal = input.InitialLocal
	itinerary.StartDay = input.StartDay
	itinerary.EndDay = input.EndDay
	itinerary.StartTime = input.StartTime
	itinerary.EndTime = input.EndTime
	itinerary.UserId = input.UserId

	newItinerary, err := s.repository.Save(itinerary)
	if err != nil {
		return newItinerary, err
	}

	return newItinerary, nil
}

func (s *service) Update(input UpdateItineraryInput) (Itinerary, error) {
	itinerary, err := s.repository.FindByID(input.ID)
	if err != nil {
		return itinerary, err
	}
	itinerary.InitialLocal = input.InitialLocal
	itinerary.StartDay = input.StartDay
	itinerary.EndDay = input.EndDay
	itinerary.StartTime = input.StartTime
	itinerary.EndTime = input.EndTime
	itinerary.UserId = input.UserId

	updatedItinerary, err := s.repository.Update(itinerary)
	if err != nil {
		return updatedItinerary, err
	}

	return updatedItinerary, nil
}

func (s *service) Delete(id int) (Itinerary, error) {
	deleteItinerary, err := s.repository.Delete(id)
	if err != nil {
		return deleteItinerary, err
	}

	return deleteItinerary, nil
}

// timeline

func (s *service) SaveTimeline(input CreateTimelineInput) (Timeline, error) {
	itinerary := Timeline{}
	itinerary.ItineraryId = input.ItineraryId
	itinerary.Time = input.Time
	itinerary.Title = input.Title
	itinerary.Description = input.Description
	itinerary.Latitude = input.Latitude
	itinerary.Longitude = input.Longitude

	newItinerary, err := s.repository.SaveTimeline(itinerary)
	if err != nil {
		return newItinerary, err
	}

	return newItinerary, nil
}

func (s *service) UpdateTimeline(input UpdateTimelineInput) (Timeline, error) {
	itinerary, err := s.repository.FindByIDTimeline(input.ID)
	if err != nil {
		return itinerary, err
	}
	itinerary.ItineraryId = input.ItineraryId
	itinerary.Time = input.Time
	itinerary.Title = input.Title
	itinerary.Description = input.Description
	itinerary.Latitude = input.Latitude
	itinerary.Longitude = input.Longitude

	updatedItinerary, err := s.repository.UpdateTimeline(itinerary)
	if err != nil {
		return updatedItinerary, err
	}

	return updatedItinerary, nil
}

func (s *service) DeleteTimeline(id int) (Timeline, error) {
	deleteItinerary, err := s.repository.DeleteTimeline(id)
	if err != nil {
		return deleteItinerary, err
	}

	return deleteItinerary, nil
}

func (s *service) FindByIDTimeline(id int) ([]Timeline, error) {
	deleteItinerary, err := s.repository.FindAllByIDTimeline(id)
	if err != nil {
		return deleteItinerary, err
	}

	return deleteItinerary, nil
}
