package seeder

import "crud/model"

var (
	Users = []*model.User{
		{Firstname: "Szymon", Lastname: "Nowak", Age: 22},
		{Firstname: "Jan", Lastname: "Kowalski", Age: 31},
		{Firstname: "Chuck", Lastname: "Norris", Age: 18},
		{Firstname: "Andrzej", Lastname: "Duda", Age: 41},
	}
)
