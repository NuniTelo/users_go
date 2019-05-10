package database

import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// Connect connects to mongodb
func Connect() {
	mongo, err := mgo.ParseURL("mongodb://localhost")
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		log.Fatal("Database connection failed: ", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})

	log.Println("Connected to", "mongodb://localhost")
	Session = s
	Mongo = mongo
}
