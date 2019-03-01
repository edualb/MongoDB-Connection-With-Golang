package main

import (
	"fmt"
	"log"
	"time"

	"./DAO"
)

// instead of `c := connection.ConnectData{Server: "localhost:27017", Database: "CRUDGolang"}`, You can use:
// c := new(connection.ConnectData)
// c.Server = "localhost:27017" <- set your Server
// c.Database = "CRUDGolang" <- set your Database
func init() {
	c := dao.ConnectData{Server: "localhost:27017", Database: "CRUDGolang"}
	_, err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create a Profile
	/*profile := new(dao.Profile)
	profile.Name = "Edualb test2"
	profile.Phone = "+55 55 955555555"
	profile.Date = time.Now()
	profile.Create()*/

	// Get All profile
	var ProfileDAO = dao.Profile{}
	profiles, _ := ProfileDAO.GetAll()
	fmt.Println("Get All:")
	for _, p := range profiles {
		fmt.Println(p)
	}

	// Get an only Profile
	profileOne, _ := ProfileDAO.GetByID("5c787a955fcd7a83b7bc8f08")
	fmt.Println("Get a only Profile:")
	fmt.Println(profileOne)

	// Update profile
	profileUpdate := new(dao.Profile)
	profileUpdate.Name = "Edualb updated2"
	profileUpdate.Phone = "+55 55 955555555"
	profileUpdate.Date = time.Now()

	// You can use profile.Update(profileUpdate) instead of profileUpdate.UpdateByID(id)
	profileUpdate.UpdateByID("5c787a955fcd7a83b7bc8f08")
	profileOne, _ = ProfileDAO.GetByID("5c787a955fcd7a83b7bc8f08")
	fmt.Println("Updated profile")
	fmt.Println(profileOne)
}
