package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
