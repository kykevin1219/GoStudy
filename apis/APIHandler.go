package apis

import (
	"fmt"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/kykevin1219/assignment/apis/netapis"
	accesslog "github.com/kykevin1219/assignment/logger"
)

type logger struct {
}

// Log ...
func (l logger) Log(record accesslog.LogRecord) {
	fmt.Printf(
		"%s Host:%s Uri:%s Method:%s Status:%d Proto:%s From:%s, Size:%d \n",
		record.Time.Format("2006-01-02T15:04:05-07:00"),
		record.Host,
		record.URI,
		record.Method,
		record.Status,
		record.Protocol,
		record.From,
		record.Size,
	)
}

// APIHandler routes apis
func APIHandler() {
	m := pat.New()
	m.Get("/ckbapiv2/network", netapis.GetNetworkInfo)
	m.Put("/ckbapiv2/network/dns", netapis.SetDNS)
	m.Put("/ckbapiv2/network/main", netapis.SetMainInterface)
	m.Put("/ckbapiv2/network/sub", netapis.SetSubInterface)
	http.Handle("/", m)
	l := logger{}
	http.ListenAndServe(":8000", accesslog.NewLoggingHandler(m, l))
}

// API rate limit
// Authentication
// Custom logger
