package models

type Project struct {
	// gorm.Model
	PCode       string `json:"pCode" gorm:"primaryKey"`
	ProjectName string `json:"projectName"`
	Dcode       string `json:"dCode"`
	Status      string `json:"status"`
	OwnerName   string `json:"ownerName"`
}
