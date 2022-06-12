package tourist

type TouristSite struct {
	ID           int
	CategoryID   int
	Name         string
	Description  string
	Address      string
	Email        string
	Phone        string
	Website      string
	Latitude     string
	Longitude    string
	LinkVideo    string
	ImagePrimary string
}

type Gallery struct {
	ID        int
	TouristID int
	Avatar    string
}
