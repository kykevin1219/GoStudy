package types

//Configuration ...
type Configuration struct {
	AddrSource   string `json:"addrSource"`
	Address      string `json:"address"`
	AllowHotplug bool   `json:"allowHotplug"`
	Auto         bool   `json:"auto"`
	Gateway      string `json:"gateway"`
	Netmask      string `json:"netmask"`
}
