package main

import (
	"./connection"
)

// instead of `c := connection.ConnectData{Server: "localhost:27017", Database: "CRUDGolang"}`, You can use:
// c := new(connection.ConnectData)
// c.Server = "localhost:27017" <- set your Server
// c.Database = "CRUDGolang" <- set your Database
func main() {
	c := connection.ConnectData{Server: "localhost:27017", Database: "CRUDGolang"}
	c.Connect()
}
