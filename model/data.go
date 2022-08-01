package model

type User struct {
	Id        int    `json:"id" form:"id" query:"id"`
	Firstname string `json:"firstname" form:"firstname" query:"firstname"`
	Lastname  string `json:"lastname" form:"lastname" query:"lastname"`
	Age       int    `json:"age" form:"age" query:"age"`
}

var (
	Users = []*User{
		{Id: 1, Firstname: "Szymon", Lastname: "Nowak", Age: 22},
		{Id: 2, Firstname: "Jan", Lastname: "Kowalski", Age: 31},
		{Id: 3, Firstname: "Chuck", Lastname: "Norris", Age: 18},
		{Id: 4, Firstname: "Andrzej", Lastname: "Duda", Age: 41},
	}
	MaxID int = 4
)

func FindUser(id int) (int, bool) {

	for i, v := range Users {
		if v.Id == id {
			return i, true
		}
	}

	return 0, false
}

func RemoveUser(arr []*User, id int) []*User {
	i, ok := FindUser(id)
	if !ok {
		return arr
	}
	return append(arr[:i], arr[i+1:]...)
}
