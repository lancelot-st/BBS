package model

type User struct {
	UserID   int64  `gorm:"unique;" json:"user_id" form:"Uuid" sql:"user_id" `
	Username string `gorm:"size:32;unique;" json:"username" form:"Name" sql:"username" validate:"required" `
	Password string `json:"password" form:"Password" sql:"password" validate:"required,max=16,min=6" `
	Email    string `gorm:"size:128" json:"email" form:"Email" sql:"email" validate:"required,email" `
}
