package connection

import (
	"fmt"

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
func (c *ConnectData) Connect() error {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		return fmt.Errorf("Error to connect on the server: %v", err)
	}
	Db = session.DB(c.Database)
	// This Check if I'm Connected in the right Database
	names, err := Db.CollectionNames()
	if err != nil {
		return fmt.Errorf("Error on collect names: %v", err)
	}
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}
