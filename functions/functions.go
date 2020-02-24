package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kykevin1219/assignment/types"
)

var network = types.Network{
	DNSServers: types.DNSServers{
		Interfaces: []string{"8.8.8.8", "8.8.4.4"},
		Resolv:     []string{"168.126.63.2", "168.126.63.1", "8.8.8.8", "8.8.4.4"},
		UpdatedAt:  1554301475,
	},
	NetworkInterfacesInfo: types.NetworkInterfacesInfo{
		Emergency: types.NetworkInterface{
			Configuration: types.Configuration{
				AddrSource:   "static",
				Address:      "192.168.62.182",
				AllowHotplug: false,
				Auto:         false,
				Gateway:      "",
				Netmask:      "255.255.0.0",
			},
			Current: types.Current{
				Address:   "",
				Broadcast: "",
				IsUp:      false,
				Netmask:   "",
			},
			Description: "",
			Name:        "eth0:1",
		},
		Gateway: "192.168.0.1",
		Main: types.NetworkInterface{
			Configuration: types.Configuration{
				AddrSource:   "dhcp",
				Address:      "192.168.0.2",
				AllowHotplug: true,
				Auto:         true,
				Gateway:      "192.168.0.1",
				Netmask:      "255.255.0.0",
			},
			Current: types.Current{
				Address:   "192.168.0.2",
				Broadcast: "192.168.0.255",
				IsUp:      true,
				Netmask:   "255.255.255.0",
			},
			Description: "",
			Name:        "eth0",
		},
		Sub: types.NetworkInterface{
			Configuration: types.Configuration{
				AddrSource:   "static",
				Address:      "10.10.20.10",
				AllowHotplug: false,
				Auto:         false,
				Gateway:      "",
				Netmask:      "255.255.255.0",
			},
			Current: types.Current{
				Address:   "",
				Broadcast: "",
				IsUp:      false,
				Netmask:   "",
			},
			Description: "",
			Name:        "eth0:2",
		},
	},
}

var errNoKey = errors.New("key error")

//GetNetworkInfo shows Network Information
func GetNetworkInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(network)
}

//SetDNS updates DNSList
func SetDNS(w http.ResponseWriter, r *http.Request) {
	var body map[string][]string
	var err = json.NewDecoder(r.Body).Decode(&body)
	if _, exist := body["dnsList"]; !exist {
		err = errNoKey
	}
	if err != nil {
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("error message"))
		fmt.Println(err)
	} else {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("SUCCESS"))
	}
}

//SetMainInterface updates MainInterface
func SetMainInterface(w http.ResponseWriter, r *http.Request) {
	var body map[string]string
	var err = json.NewDecoder(r.Body).Decode(&body)
	//디코드가 안되었을경우 에러핸들링
	var keys = []string{"addrSource", "address", "gateway", "netmask"}
	for _, key := range keys {
		if _, exist := body[key]; !exist {
			err = errNoKey
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("SUCCESS"))
}

//SetSubInterface updates MainInterface
func SetSubInterface(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	var err = json.NewDecoder(r.Body).Decode(&body)
	var keys = []string{"auto", "address", "netmask"}
	for _, key := range keys {
		if _, exist := body[key]; !exist {
			err = errNoKey
		}
	}
	if err != nil {
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("error message"))
	} else {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte("SUCCESS"))
	}
}
