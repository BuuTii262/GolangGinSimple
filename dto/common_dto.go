package dto

type DeleteByIdDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}
