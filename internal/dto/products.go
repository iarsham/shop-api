package dto

type ProductRequest struct {
	Name        string   `json:"name" example:"Table" binding:"required,min=1,max=75"`
	Description string   `json:"description" example:"Contrary to popular belief, Lorem Ipsum is not simply random text" binding:"required,min=1,max=300"`
	Price       float64  `json:"price" example:"19.99" binding:"required"`
	Stock       int      `json:"stock" example:"12" binding:"required"`
	Weight      int      `json:"weight" example:"5" binding:"required"`
	Tags        []string `json:"tags" example:"ai,ml" binding:"omitempty"`
}
