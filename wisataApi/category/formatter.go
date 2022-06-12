package category

type CategoryFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

func FormatCategory(category Category) CategoryFormatter {
	CategoryFormatter := CategoryFormatter{}
	CategoryFormatter.ID = category.ID
	CategoryFormatter.Name = category.Name
	CategoryFormatter.Avatar = category.Avatar

	return CategoryFormatter
}

func FormatCategorys(campaigns []Category) []CategoryFormatter {
	campaignsFormatter := []CategoryFormatter{}

	for _, campaign := range campaigns {
		CategoryFormatter := FormatCategory(campaign)
		campaignsFormatter = append(campaignsFormatter, CategoryFormatter)
	}

	return campaignsFormatter
}
