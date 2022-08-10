package model

type User struct {
	Id        int    `json:"id" form:"id" query:"id" gorm:"primaryKey;autoIncrement;not null;<-:create"`
	Firstname string `json:"firstname" form:"firstname" query:"firstname" gorm:"size:50"`
	Lastname  string `json:"lastname" form:"lastname" query:"lastname" gorm:"size:50"`
	Age       int    `json:"age" form:"age" query:"age"`
}
