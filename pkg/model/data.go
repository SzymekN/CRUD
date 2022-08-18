package model

// swagger:model User
type User struct {
	// Id of the user
	// in: int64
	// required: false
	Id int `json:"id" form:"id" query:"id" gorm:"primaryKey;autoIncrement;not null;<-:create"`
	// Firstname of the user
	// in: string
	// required: true
	// minimum length: 3
	// maximum length: 30
	Firstname string `json:"firstname" form:"firstname" query:"firstname" gorm:"size:50"`
	// Lastname of the user
	// in: string
	// required: true
	// minimum length: 3
	// maximum length: 30
	Lastname string `json:"lastname" form:"lastname" query:"lastname" gorm:"size:50"`
	// Age of the user
	// in: int64
	// required: true
	// minimum: 18
	// maximum: 99
	Age int `json:"age" form:"age" query:"age"`
}

// swagger:model GenericError
type GenericError struct {
	// Response message
	// in: string
	// required: true
	Message string `json:"message"`
}

// swagger:model GenericMessage
type GenericMessage struct {
	// Response error with message
	// in: string
	// required: true
	Message string `json:"message"`
}
