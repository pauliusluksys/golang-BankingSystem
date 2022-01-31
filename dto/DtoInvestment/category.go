package DtoInvestment

import "gorm.io/gorm"

type CategoryRequest struct {
	gorm.Model
	Name string
}
type NewCategoryRequest struct {
	Name string `json:"name"`
}
type NewCategoryResponse struct {
	ID uint `json:"id"`
}
