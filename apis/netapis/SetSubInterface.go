package netapis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

//SetSubInterfaceBody ...
type SetSubInterfaceBody struct {
	Auto    bool   `json:"auto"`
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}

//SetSubInterface updates MainInterface
func SetSubInterface(w http.ResponseWriter, r *http.Request) {
	//DECODE
	var reqbody SetSubInterfaceBody
	err := json.NewDecoder(r.Body).Decode(&reqbody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//CHECK the content of Address, Netmask
	IPs := [2]string{reqbody.Address, reqbody.Netmask}
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
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
	w.WriteHeader(200)
	w.Write([]byte("SUCCESS"))
	return
}
