package tourist

type CreateTouristInput struct {
	CategoryID   int    `json:"category_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Website      string `json:"website" binding:"required"`
	Latitude     string `json:"latitude" binding:"required"`
	Longitude    string `json:"longitude" binding:"required"`
	LinkVideo    string `json:"link_video" binding:"required"`
	ImagePrimary string `json:"image_primary" binding:"required"`
}

type UpdateTouristInput struct {
	ID           int    `json:"tourist_id" binding:"required"`
	CategoryID   int    `json:"category_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Website      string `json:"website" binding:"required"`
	Latitude     string `json:"latitude" binding:"required"`
	Longitude    string `json:"longitude" binding:"required"`
	LinkVideo    string `json:"link_video" binding:"required"`
	ImagePrimary string `json:"image_primary" binding:"required"`
}

type DeleteTouristInput struct {
	ID int `json:"tourist_id" binding:"required"`
}

type CreateGelleryTouristInput struct {
	TouristID int    `json:"tourist_id" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

type DeleteGalleryTouristInput struct {
	ID int `json:"gallery_id" binding:"required"`
}

type FindAllGalleryTouristByID struct {
	ID int `json:"tourist_id"`
}

type AllTouristByCategory struct {
	CategoryID int `json:"category_id" binding:"required"`
}

type Search struct {
	Name string `json:"name" binding:"required"`
}
