package types

//NetworkInterfacesInfo ...
type NetworkInterfacesInfo struct {
	Emergency NetworkInterface `json:"emergency"`
	Gateway   string           `json:"gateway"`
	Main      NetworkInterface `json:"main"`
	Sub       NetworkInterface `json:"sub"`
}
