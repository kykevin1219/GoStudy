package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/kykevin1219/assignment/functions"
)

func main() {
	m := pat.New()
	m.Get("/ckbapiv2/network", functions.GetNetworkInfo)
	m.Put("/ckbapiv2/network/dns", functions.SetDNS)
	m.Put("/ckbapiv2/network/main", functions.SetMainInterface)
	m.Put("/ckbapiv2/network/sub", functions.SetSubInterface)

	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
