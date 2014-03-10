package romp

import (
	"github.com/DataPlatform/romp/models"
	"github.com/DataPlatform/romp/resources"
	"github.com/gorilla/mux"
	"net/http"
	"time"

//	"github.com/gorilla/sessions"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	//	vars := mux.Vars(r)
}

func init_router() *mux.Router {
	var router = new(mux.Router)
	router.HandleFunc("/", homeHandler).Name("home")
	reportEndpoint := resources.HTTPEndpoint{resources.ReportResource{models.GetCollection("report")}}

	router.HandleFunc("/report", reportEndpoint.Handle).Name("report")
	return router
}

func Init_server() *http.Server {
	router := init_router()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}
