package netapis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

//SetMainInterfaceBody ...
type SetMainInterfaceBody struct {
	AddrSource string `json:"addrSource"`
	Address    string `json:"address"`
	Gateway    string `json:"gateway"`
	Netmask    string `json:"netmask"`
}

//SetMainInterface updates MainInterface
func SetMainInterface(w http.ResponseWriter, r *http.Request) {
	//DECODE
	var reqbody SetMainInterfaceBody
	err := json.NewDecoder(r.Body).Decode(&reqbody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//CHECK AddrSource
	if reqbody.AddrSource != "dhcp" && reqbody.AddrSource != "static" {
		http.Error(w, "Wrong AddrSource", http.StatusBadRequest)
		return
	}

	//CHECK the Content of Adress, Gateway, Netmask
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	var IPs = [3]string{reqbody.Address, reqbody.Gateway, reqbody.Netmask}
	for _, stringIP := range IPs {
		IP := re.FindString(stringIP)
		if IP == "" {
			err = fmt.Errorf("%s is not a proper IP address", stringIP)
			break
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Done
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}
