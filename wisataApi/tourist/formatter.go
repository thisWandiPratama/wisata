package tourist

type TouristFormatter struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"category_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Website      string `json:"website"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	LinkVideo    string `json:"link_video"`
	ImagePrimary string `json:"image_primary"`
}

func FormatTourist(tourist TouristSite) TouristFormatter {
	TouristFormatter := TouristFormatter{}
	TouristFormatter.ID = tourist.ID
	TouristFormatter.CategoryID = tourist.CategoryID
	TouristFormatter.Name = tourist.Name
	TouristFormatter.Description = tourist.Description
	TouristFormatter.Address = tourist.Address
	TouristFormatter.Email = tourist.Email
	TouristFormatter.Phone = tourist.Phone
	TouristFormatter.Website = tourist.Website
	TouristFormatter.Latitude = tourist.Latitude
	TouristFormatter.Longitude = tourist.Longitude
	TouristFormatter.LinkVideo = tourist.LinkVideo
	TouristFormatter.ImagePrimary = tourist.ImagePrimary

	return TouristFormatter
}

func FormatTourists(tourists []TouristSite) []TouristFormatter {
	touristsFormatter := []TouristFormatter{}

	for _, tourists := range tourists {
		TouristFormatter := FormatTourist(tourists)
		touristsFormatter = append(touristsFormatter, TouristFormatter)
	}

	return touristsFormatter
}

type GalleryTouristFormatter struct {
	ID        int    `json:"id"`
	TouristID int    `json:"tourist_id"`
	Avatar    string `json:"avatar"`
}

func FormatGalleryTourist(tourist Gallery) GalleryTouristFormatter {
	TouristFormatter := GalleryTouristFormatter{}
	TouristFormatter.ID = tourist.ID
	TouristFormatter.TouristID = tourist.TouristID
	TouristFormatter.Avatar = tourist.Avatar

	return TouristFormatter
}

func FormatGalleryTourists(tourists []Gallery) []GalleryTouristFormatter {
	touristsFormatter := []GalleryTouristFormatter{}

	for _, tourists := range tourists {
		TouristFormatter := FormatGalleryTourist(tourists)
		touristsFormatter = append(touristsFormatter, TouristFormatter)
	}

	return touristsFormatter
}
