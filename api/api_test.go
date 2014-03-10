package romp

import (
	"fmt"
	// "github.com/DataPlatform/romp/models"
	// "github.com/DataPlatform/romp/resources"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func init_test_server() *httptest.Server {
	router := init_router()
	ts := httptest.NewServer(router)
	return ts
}

// Basic tests
func TestServerUp(t *testing.T) {
	ts := init_test_server()
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello from server:", greeting)
}

func TestAddReport(t *testing.T) {
	ts := init_test_server()
	defer ts.Close()
	report := url.Values{}
	report.Add("template", "Amazing")
	report.Add("queries", "1,2")
	res, err := http.PostForm(ts.URL+"/report", report)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

}
