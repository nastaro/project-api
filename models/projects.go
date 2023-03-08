package models

type Project struct {
	// gorm.Model
	PCode       string `json:"pCode" gorm:"primaryKey"`
	ProjectName string `json:"projectName" binding:"required"`
	Dcode       string `json:"dCode" binding:"required"`
	Status      string `json:"status"`
	OwnerName   string `json:"ownerName" binding:"required"`
}
