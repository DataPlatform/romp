package main

import (
	"github.com/DataPlatform/romp/api"
)

func main() {
	s := romp.Init_server()
	s.ListenAndServe()
}
