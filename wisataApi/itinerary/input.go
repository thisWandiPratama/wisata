package itinerary

type CreateItineraryInput struct {
	InitialLocal string `json:"intial_local" binding:"required"`
	StartDay     string `json:"start_day" binding:"required"`
	EndDay       string `json:"end_day" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"`
	EndTime      string `json:"end_time" binding:"required"`
	UserId       int    `json:"user_id" binding:"required"`
}

type CreateTimelineInput struct {
	ItineraryId int    `json:"itinerary_id" binding:"required"`
	Time        string `json:"time" binding:"required"`        //day
	Title       string `json:"title" binding:"required"`       //time
	Description string `json:"description" binding:"required"` //name
	Latitude    string `json:"latitude" binding:"required"`
	Longitude   string `json:"longitude" binding:"required"`
	Jarak       int    `json:"jarak" binding:"required"`
}

type UpdateTimelineInput struct {
	ID          int    `json:"timeline_id" binding:"required"`
	ItineraryId int    `json:"itinerary_id" binding:"required"`
	Time        string `json:"time" binding:"required"`        //day
	Title       string `json:"title" binding:"required"`       //time
	Description string `json:"description" binding:"required"` //name
	Latitude    string `json:"latitude" binding:"required"`
	Longitude   string `json:"longitude" binding:"required"`
	Jarak       int    `json:"jarak" binding:"required"`
}

type UpdateItineraryInput struct {
	ID           int    `json:"itinerary_id" binding:"required"`
	InitialLocal string `json:"intial_local" binding:"required"`
	StartDay     string `json:"start_day" binding:"required"`
	EndDay       string `json:"end_day" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"`
	EndTime      string `json:"end_time" binding:"required"`
	UserId       int    `json:"user_id" binding:"required"`
}

type DeleteItineraryInput struct {
	ID int `json:"itinerary_id" binding:"required"`
}

type FindAllItineraryByUser struct {
	ID int `json:"user_id" binding:"required"`
}
