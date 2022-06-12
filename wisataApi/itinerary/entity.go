package itinerary

type Itinerary struct {
	ID           int
	InitialLocal string
	StartDay     string
	EndDay       string
	StartTime    string
	EndTime      string
	UserId       int
}

type Timeline struct {
	ID          int
	ItineraryId int
	Time        string //day
	Title       string //time
	Description string //name
	Latitude    string
	Longitude   string
}
