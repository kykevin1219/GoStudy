package apis

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/kykevin1219/assignment/apis/netapis"
)

//APIHandler routes apis
func APIHandler() {
	m := pat.New()
	m.Get("/ckbapiv2/network", netapis.GetNetworkInfo)
	m.Put("/ckbapiv2/network/dns", netapis.SetDNS)
	m.Put("/ckbapiv2/network/main", netapis.SetMainInterface)
	m.Put("/ckbapiv2/network/sub", netapis.SetSubInterface)
	http.Handle("/", m)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
