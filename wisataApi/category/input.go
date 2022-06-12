package category

type CreateCategoryInput struct {
	Name   string `json:"name" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

type UpdateCategoryInput struct {
	ID     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Avatar string `json:"avatar" binding:"required"`
}

type DeleteCategoryInput struct {
	ID int `json:"id_category" binding:"required"`
}
