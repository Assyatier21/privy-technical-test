package models

type Cake struct {
	Id          int     `json:"id" form:"id"`
	Title       string  `json:"title" form:"title"`
	Description string  `json:"description" form:"description"`
	Rating      float32 `json:"rating" form:"rating"`
	Image       string  `json:"image" form:"image"`
	CreatedAt   string  `json:"created_at" form:"created_at"`
	UpdatedAt   string  `json:"updated_at" form:"updated_at"`
}
