package model

type ShopeeCateData struct {
	Data DataCat `bson:"data"`
}

type DataCat struct {
	CategoryList []CategoryList `json:"category_list"`
}

type CategoryList struct {
	Catid       int    `json:"catid"`
	ParentCatid int    `json:"parent_catid"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Image       string `json:"image"`
	Level       int    `json:"level"`
	Children    []struct {
		Catid       int         `json:"catid"`
		ParentCatid int         `json:"parent_catid"`
		Name        string      `json:"name"`
		DisplayName string      `json:"display_name"`
		Image       string      `json:"image"`
		Level       int         `json:"level"`
		Children    interface{} `json:"children"`
	}
}
