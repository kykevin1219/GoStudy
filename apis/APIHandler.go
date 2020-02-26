package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/kykevin1219/assignment/apis/netapis"
	accesslog "github.com/mash/go-accesslog"
)

type logger struct {
}

// Log ...
func (l logger) Log(record accesslog.LogRecord) {
	fmt.Printf(
		"%s Host:%s Uri:%s Method:%s Status:%d Proto:%s From:%s, Size:%d \n",
		time.Now().Format("2006-01-02T15:04:05-07:00"),
		record.Host,
		record.Uri,
		record.Method,
		record.Status,
		record.Protocol,
		record.Ip,
		record.Size,
		// size , who,
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
