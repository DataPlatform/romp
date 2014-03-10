package resources

import (
	"encoding/json"
	// "github.com/DataPlatform/romp/models"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

//A type for exchanging slightly structured data
type ResourceData struct{}

//Implement this interface for each model you wish to expose as a resource
type Resource interface {
	Put(string, url.Values) ResourceData
	Post(url.Values) ResourceData
	Get(string) ResourceData
	Delete(string) ResourceData
}

// HTTPEndpoint wraps a Resource and handles unpacking
// of the http.Request into typed data
type HTTPEndpoint struct {
	R Resource
}

func (end HTTPEndpoint) Handle(w http.ResponseWriter, r *http.Request) {
	//pull named parameters from the routed url
	vars := mux.Vars(r)
	r.ParseForm()
	//loot at the http verb and call the appropriate function on the resource
	fmt.Println(r.Method + " to " + r.URL.String())
	if r.Method == "GET" {
		data, err := json.Marshal(end.R.Get(vars["id"]))
		if err == nil {
			w.Write(data)
		}
	}
	if r.Method == "POST" {
		fmt.Println("Data: ", r.Form)
		data, err := json.Marshal(end.R.Post(r.Form))
		if err == nil {
			w.Write(data)
		}
	}
	if r.Method == "PUT" {
		fmt.Println("Data: ", r.Form)
		data, err := json.Marshal(end.R.Put(vars["id"], r.Form))
		if err == nil {
			w.Write(data)
		}
	}
	if r.Method == "DELETE" {
		data, err := json.Marshal(end.R.Delete(vars["id"]))
		if err == nil {
			w.Write(data)
		}
	}

}
