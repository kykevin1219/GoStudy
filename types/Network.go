package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Network ...
type Network struct {
	DNSServers            DNSServers            `json:"dnsServers"`
	NetworkInterfacesInfo NetworkInterfacesInfo `json:"networkInterfacesInfo"`
}

//InitData ...
func (n *Network) InitData() {
	jsonfile, err := os.Open("network.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonfile)
	json.Unmarshal(byteValue, &n)
}
