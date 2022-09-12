package seeder

import (
	"fmt"
	"log"

	"github.com/SzymekN/CRUD/pkg/auth"
	"github.com/SzymekN/CRUD/pkg/model"
	"github.com/SzymekN/CRUD/pkg/storage"
	"github.com/gocql/gocql"

	"github.com/jinzhu/gorm"
)

type SeedPG struct {
	Name string
	Run  func(*gorm.DB) error
}
type SeedCASS struct {
	Name string
	Run  func(*gocql.Session) error
}

func AllPG() []SeedPG {
	return []SeedPG{
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

func AllCASS() []SeedCASS {
	return []SeedCASS{
		{
			Name: "Create users",
			Run: func(cas *gocql.Session) error {
				var err error

				for _, u := range Users {
					err = cas.Query(`Insert into userapi.users(id, firstname,lastname, age) values (?,?,?,?)`, u.Id, u.Firstname, u.Lastname, u.Age).Exec()
					if err != nil {
						return err
					}
				}

				return nil
			},
		},
		{
			Name: "Create operators",
			Run: func(cas *gocql.Session) error {
				var err error

				for _, o := range Operators {
					// tu może nie działać coś
					pwd, _ := auth.GeneratehashPassword(o.Password)
					err = cas.Query(`Insert into userapi.operators(username, email,password, role) values (?,?,?,?)`, o.Username, o.Email, pwd, o.Role).Exec()
					if err != nil {
						return err
					}
				}

				return nil
			},
		},
	}
}

func CreateAndSeed() {
	CreateAndSeedPG(storage.GetDBInstance())
	CreateAndSeedCASS(storage.GetCassandraInstance())
}

func CreateAndSeedPG(db *gorm.DB) {

	tableExists := db.HasTable("users")
	if tableExists {

		err := db.DropTable("users").Error

		if err != nil {
			log.Fatalf("PG: Dropping table users failed with error: %s", err)
		} else {
			fmt.Println("PG: SUCCESFULLY dropped table users!")
		}

	}

	err := db.Table("users").AutoMigrate(&model.User{}).Error

	if err != nil {
		log.Fatalf("PG: Creating table users failed with error: %s", err)
	} else {
		fmt.Println("PG: SUCCESFULLY created table users!")
	}

	for _, seed := range AllPG() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("PG: Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
	fmt.Println("PG: SUCCESFULLY seeded table database!")

}
func CreateAndSeedCASS(cas *gocql.Session) {

	err := cas.Query("CREATE KEYSPACE IF NOT EXISTS userapi WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };").Exec()
	if err != nil {
		log.Fatalf("CASS: Creating table users failed with error: %s", err)
	} else {
		fmt.Println("CASS: SUCCESFULLY dropped table users!")
	}

	err = cas.Query(`CREATE TABLE IF NOT EXISTS userapi.users (
		id int PRIMARY KEY,
		firstname text ,
		lastname text,
		age int
		);`).Exec()

	if err != nil {
		log.Fatalf("CAS: Creating table users failed with error: %s", err)
	} else {
		fmt.Println("CAS: SUCCESFULLY created table users!")
	}

	err = cas.Query(`CREATE TABLE IF NOT EXISTS userapi.operators (
		username text PRIMARY KEY,
		email text,
		password text,
		role text
		);`).Exec()

	if err != nil {
		log.Fatalf("CAS: Creating table operators failed with error: %s", err)
	} else {
		fmt.Println("CAS: SUCCESFULLY created table operators!")
	}

	for _, seed := range AllCASS() {
		if err := seed.Run(cas); err != nil {
			log.Fatalf("CAS: Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
	fmt.Println("CAS: SUCCESFULLY seeded table database!")

}
