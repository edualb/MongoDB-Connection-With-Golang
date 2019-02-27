package dao

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
func (c *ConnectData) Connect() (*mgo.Session, error) {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		return nil, fmt.Errorf("Error to connect on the server: %v", err)
	}
	session.SetMode(mgo.Monotonic, true)
	Db = session.DB(c.Database)
	return session, nil
}
