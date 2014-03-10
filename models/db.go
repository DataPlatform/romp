package models

import (
	"labix.org/v2/mgo"
)

func GetSession() mgo.Session {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// // Optional. Switch the session to a monotonic behavior.
	// session.SetMode(mgo.Monotonic, true)
	return *session
}

func GetCollection(db string) *mgo.Collection {
	session := GetSession()
	c := session.DB("romp").C(db)
	return c
}
