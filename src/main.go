package main

import (
	"log"

	"./DAO"
	"./connection"
)

// instead of `c := connection.ConnectData{Server: "localhost:27017", Database: "CRUDGolang"}`, You can use:
// c := new(connection.ConnectData)
// c.Server = "localhost:27017" <- set your Server
// c.Database = "CRUDGolang" <- set your Database
func init() {
	c := connection.ConnectData{Server: "localhost:27017", Database: "CRUDGolang"}
	_, err := c.Connect()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create a Profile
	profile := new(DAO.ProfileDAO)
	profile.Name = "Edualb"
	profile.Phone = "+55 55 955555555"
	profile.Create()
}
