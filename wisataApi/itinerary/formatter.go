package itinerary

type ItineraryFormatter struct {
	ID           int    `json:"id"`
	InitialLocal string `json:"intial_local"`
	StartDay     string `json:"start_day"`
	EndDay       string `json:"end_day"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	UserId       int    `json:"user_id"`
}

func FormatItinerary(itinerary Itinerary) ItineraryFormatter {
	ItineraryFormatter := ItineraryFormatter{}
	ItineraryFormatter.ID = itinerary.ID
	ItineraryFormatter.InitialLocal = itinerary.InitialLocal
	ItineraryFormatter.StartDay = itinerary.StartDay
	ItineraryFormatter.EndDay = itinerary.EndDay
	ItineraryFormatter.StartTime = itinerary.StartTime
	ItineraryFormatter.EndTime = itinerary.EndTime
	ItineraryFormatter.UserId = itinerary.UserId

	return ItineraryFormatter
}

func FormatItinerarys(itinerarys []Itinerary) []ItineraryFormatter {
	itinerarysFormatter := []ItineraryFormatter{}

	for _, itinerary := range itinerarys {
		ItineraryFormatter := FormatItinerary(itinerary)
		itinerarysFormatter = append(itinerarysFormatter, ItineraryFormatter)
	}

	return itinerarysFormatter
}

type TimelineFormatter struct {
	ID          int    `json:"id"`
	ItineraryId int    `json:"itinerary_id"`
	Time        string `json:"time"`        //day
	Title       string `json:"title"`       //time
	Description string `json:"description"` //name
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

func FormatTimeline(itinerary Timeline) TimelineFormatter {
	ItineraryFormatter := TimelineFormatter{}
	ItineraryFormatter.ID = itinerary.ID
	ItineraryFormatter.ItineraryId = itinerary.ItineraryId
	ItineraryFormatter.Time = itinerary.Time
	ItineraryFormatter.Title = itinerary.Title
	ItineraryFormatter.Description = itinerary.Description
	ItineraryFormatter.Latitude = itinerary.Latitude
	ItineraryFormatter.Longitude = itinerary.Longitude

	return ItineraryFormatter
}

func FormatTimelines(itinerarys []Timeline) []TimelineFormatter {
	itinerarysFormatter := []TimelineFormatter{}

	for _, itinerary := range itinerarys {
		ItineraryFormatter := FormatTimeline(itinerary)
		itinerarysFormatter = append(itinerarysFormatter, ItineraryFormatter)
	}

	return itinerarysFormatter
}
