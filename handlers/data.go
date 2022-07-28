package handlers

type User struct {
	Id         int    `json:"id" query:"id" form:"id" param:"id"`
	First_name string `json:"firstname" query:"firstname" form:"firstname" param:"firstname"`
	Last_name  string `json:"lastname" query:"lastname" form:"lastname" param:"lastname"`
	Age        int    `json:"age" query:"age" form:"age" param:"age"`
}

var (
	users = map[int]*User{
		1: {Id: 1, First_name: "Szymon", Last_name: "Nowak", Age: 22},
		2: {Id: 2, First_name: "Jan", Last_name: "Kowalski", Age: 31},
		3: {Id: 3, First_name: "Chuck", Last_name: "Norris", Age: 18},
		4: {Id: 4, First_name: "Andrzej", Last_name: "Duda", Age: 41},
	}
	maxID int = 4
)
