package resources

import (
	"github.com/DataPlatform/romp/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/url"
	"strings"
	"time"
)

type ReportResource struct {
	Db *mgo.Collection
}

func (rr ReportResource) Get(id string) ResourceData {
	return *new(ResourceData)
}
func (rr ReportResource) Post(data url.Values) ResourceData {
	rr.Db.Insert()
	return *new(ResourceData)
}
func (rr ReportResource) Put(id string, data url.Values) ResourceData {
	return *new(ResourceData)
}
func (rr ReportResource) Delete(string) ResourceData {
	return *new(ResourceData)
}

func ReportFromFormData(data url.Values) {
	r := models.Report{
		Id:        bson.NewObjectId(),
		Timestamp: time.Now(),
		Template:  data.Get("template"),
	}
	query_values := strings.Split(data.Get("queries"), ",")
	queries := make([]bson.ObjectId, len(query_values))
	for q := range query_values {
		queries = append(queries, bson.ObjectId(q))
	}
	r.Queries = queries
}
