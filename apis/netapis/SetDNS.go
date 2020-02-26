package netapis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

//SetDNSRequestBody ...
type SetDNSRequestBody struct {
	DNSList []string `json:"dnsList"`
}

//SetDNS updates DNSList
func SetDNS(w http.ResponseWriter, r *http.Request) {
	//DECODE
	var reqbody SetDNSRequestBody
	var err = json.NewDecoder(r.Body).Decode(&reqbody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//CHECK the number of DNS List
	if len(reqbody.DNSList) == 0 {
		http.Error(w, "At least one DNS needed", http.StatusBadRequest)
		return
	}

	//CHECK the content of DNS List
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	for _, DNS := range reqbody.DNSList {
		var ip string = re.FindString(DNS)
		if ip == "" {
			err = fmt.Errorf("%s is not a proper IP address", DNS)
			break
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//DONE
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
	return
}
