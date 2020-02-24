package types

type DNSServers struct {
	Interfaces []string `json:"interfaces"`
	Resolv     []string `json:"resolv"`
	UpdatedAt  int      `json:"updatedAt"`
}

type Configuration struct {
	AddrSource   string `json:"addrSource"`
	Address      string `json:"address"`
	AllowHotplug bool   `json:"allowHotplug"`
	Auto         bool   `json:"auto"`
	Gateway      string `json:"gateway"`
	Netmask      string `json:"netmask"`
}

type Current struct {
	Address   string `json:"address"`
	Broadcast string `json:"broadcast"`
	IsUp      bool   `json:"isUp"`
	Netmask   string `json:"netmask"`
}

type NetworkInterface struct {
	Configuration Configuration `json:"configuration"`
	Current       Current       `json:"current"`
	Description   string        `json:"description"`
	Name          string        `json:"name"`
}

type NetworkInterfacesInfo struct {
	Emergency NetworkInterface `json:"emergency"`
	Gateway   string           `json:"gateway"`
	Main      NetworkInterface `json:"main"`
	Sub       NetworkInterface `json:"sub"`
}
type Network struct {
	DNSServers            DNSServers            `json:"dnsServers"`
	NetworkInterfacesInfo NetworkInterfacesInfo `json:"networkInterfacesInfo"`
}
