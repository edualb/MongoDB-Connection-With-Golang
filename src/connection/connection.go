package connection

import (
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

// ConnectData is an exported type that
// contains two strings to connect
// on the database
type ConnectData struct {
	Server   string
	Database string
}

// Db variable used to get the connection from MongoDb
var Db *mgo.Database

// Connect request a connection with MongoDb
func (c *ConnectData) Connect() {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	Db = session.DB(c.Database)
	// This Check if I'm Connected in the right Database
	names, err := Db.CollectionNames()
	for _, name := range names {
		fmt.Println(name)
	}
}
