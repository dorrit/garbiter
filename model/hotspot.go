package model

type HotspotServer struct {
	ID          string
	Name        string
	Interface   string
	AddressPool string
	Profile     string
	IdleTimeout string
	Disabled    bool
	Raw         map[string]string
}

type HotspotServerSet struct {
	Name        string
	Interface   string
	AddressPool string
	Profile     string
	IdleTimeout string
	Disabled    *bool
	Extra       map[string]string
}

type HotspotUser struct {
	ID       string
	Name     string
	Profile  string
	Server   string
	Address  string
	Disabled bool
	Comment  string
	Raw      map[string]string
}

type HotspotUserSet struct {
	Name     string
	Password string
	Profile  string
	Server   string
	Address  string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}

type HotspotActive struct {
	ID         string
	User       string
	Address    string
	MACAddress string
	Uptime     string
	Server     string
	Raw        map[string]string
}
