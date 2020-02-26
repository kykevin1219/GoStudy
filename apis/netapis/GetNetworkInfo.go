package netapis

import (
	"encoding/json"
	"net/http"

	"github.com/kykevin1219/assignment/types"
)

//GetNetworkInfo shows Network Information
func GetNetworkInfo(w http.ResponseWriter, r *http.Request) {
	var network types.Network
	network.InitData()
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(network)
}
