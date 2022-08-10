package seeder

import (
	"crud/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "Create users",
			Run: func(db *gorm.DB) error {
				var err error

				for _, user := range Users {
					err = db.Create(&user).Error
					if err != nil {
						return err
					}
				}

				return nil
			},
		},
	}
}

func CreateAndSeed(db *gorm.DB) {

	if db.HasTable("users") {
		err := db.DropTable("users").Error

		if err != nil {
			log.Fatalf("Dropping table users failed with error: %s", err)
		} else {
			fmt.Println("SUCCESFULLY dropped table users!")
		}

	}

	err := db.Table("users").AutoMigrate(&model.User{}).Error

	if err != nil {
		log.Fatalf("Creating table users failed with error: %s", err)
	} else {
		fmt.Println("SUCCESFULLY created table users!")
	}

	for _, seed := range All() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
	fmt.Println("SUCCESFULLY seeded table database!")

}
