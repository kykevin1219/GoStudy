package types

//Current ...
type Current struct {
	Address   string `json:"address"`
	Broadcast string `json:"broadcast"`
	IsUp      bool   `json:"isUp"`
	Netmask   string `json:"netmask"`
}
