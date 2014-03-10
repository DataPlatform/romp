package resources

import (
	// "github.com/DataPlatform/romp/models"
	"labix.org/v2/mgo"
	// "labix.org/v2/mgo/bson"
)

type QueryResource struct {
	Db mgo.Collection
}

func (qr QueryResource) Get(id string) ResourceData {
	return *new(ResourceData)
}
func (qr QueryResource) Post(data ResourceData) ResourceData {
	return *new(ResourceData)
}
func (qr QueryResource) Put(id string, data ResourceData) ResourceData {
	return *new(ResourceData)
}
func (qr QueryResource) Delete(string) ResourceData {
	return *new(ResourceData)
}
