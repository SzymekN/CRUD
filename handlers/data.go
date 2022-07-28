package handlers

type User struct {
	Id         int    `json:"id" query:"id" form:"id" param:"id"`
	First_name string `json:"firstname" query:"firstname" form:"firstname" param:"firstname"`
	Last_name  string `json:"lastname" query:"lastname" form:"lastname" param:"lastname"`
	Age        int    `json:"age" query:"age" form:"age" param:"age"`
}

var (
	users = []*User{
		{Id: 1, First_name: "Szymon", Last_name: "Nowak", Age: 22},
		{Id: 2, First_name: "Jan", Last_name: "Kowalski", Age: 31},
		{Id: 3, First_name: "Chuck", Last_name: "Norris", Age: 18},
		{Id: 4, First_name: "Andrzej", Last_name: "Duda", Age: 41},
	}
	maxID int = 4
)

func findUser(id int) (int, bool) {

	for i, v := range users {
		if v.Id == id {
			return i, true
		}
	}

	return 0, false
}

func removeUser(arr []*User, id int) []*User {
	i, ok := findUser(id)
	if !ok {
		return arr
	}
	return append(arr[:i], arr[i+1:]...)
}
