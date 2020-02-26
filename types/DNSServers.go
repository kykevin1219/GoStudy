package types

//DNSServers ...
type DNSServers struct {
	Interfaces []string `json:"interfaces"`
	Resolv     []string `json:"resolv"`
	UpdatedAt  int      `json:"updatedAt"`
}
