package types

//NetworkInterface ...
type NetworkInterface struct {
	Configuration Configuration `json:"configuration"`
	Current       Current       `json:"current"`
	Description   string        `json:"description"`
	Name          string        `json:"name"`
}
