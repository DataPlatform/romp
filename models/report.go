package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Report struct {
	Id        bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Timestamp time.Time       `json:"timestamp"`
	Template  string          `json:"template"`
	Queries   []bson.ObjectId `json:"queries"`
}

type ReportDocument struct {
	Report
	Queries []Query
}

func (r Report) AsDocument(collection *mgo.Collection) ReportDocument {
	rd := ReportDocument{r, make([]Query, len(r.Queries))}
	for i := range r.Queries {
		//init Query object with only the Id populated
		query := new(Query)
		err := collection.FindId(r.Queries[i]).One(&query)
		if err == nil {
			rd.Queries[i] = *query
		}
	}
	return rd
}

func (rd ReportDocument) AsRecord() Report {
	r := Report{rd.Id, rd.Timestamp, rd.Template, make([]bson.ObjectId, len(rd.Queries))}
	for i := range rd.Queries {
		r.Queries[i] = rd.Queries[i].Id
	}
	return r
}
