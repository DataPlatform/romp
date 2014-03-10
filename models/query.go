package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Query struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Query     string
	Timestamp time.Time
}
