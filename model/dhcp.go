package model

type DHCPClient struct {
	ID        string
	Interface string
	Address   string
	Gateway   string
	Status    string
	Disabled  bool
	Comment   string
	Raw       map[string]string
}

type DHCPClientSet struct {
	Interface       string
	UsePeerDNS      *bool
	UsePeerNTP      *bool
	AddDefaultRoute string
	Comment         string
	Disabled        *bool
	Extra           map[string]string
}

type DHCPServer struct {
	ID          string
	Name        string
	Interface   string
	AddressPool string
	LeaseTime   string
	Disabled    bool
	Comment     string
	Raw         map[string]string
}

type DHCPServerSet struct {
	Name        string
	Interface   string
	AddressPool string
	LeaseTime   string
	Comment     string
	Disabled    *bool
	Extra       map[string]string
}

type DHCPLease struct {
	ID         string
	Address    string
	MACAddress string
	HostName   string
	Server     string
	Status     string
	Dynamic    bool
	Disabled   bool
	Comment    string
	Raw        map[string]string
}
