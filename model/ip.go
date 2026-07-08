package model

type IPAddress struct {
	ID        string
	Address   string
	Network   string
	Interface string
	Disabled  bool
	Comment   string
	Raw       map[string]string
}

type IPAddressSet struct {
	Address   string
	Network   string
	Interface string
	Comment   string
	Disabled  *bool
	Extra     map[string]string
}

type Route struct {
	ID       string
	Dst      string
	Gateway  string
	Distance string
	Scope    string
	Target   string
	Active   bool
	Disabled bool
	Comment  string
	Raw      map[string]string
}

type RouteSet struct {
	Dst      string
	Gateway  string
	Distance string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}

type DNS struct {
	Servers             string
	DynamicServers      string
	AllowRemoteRequests bool
	CacheSize           string
	Raw                 map[string]string
}

type DNSSet struct {
	Servers             string
	AllowRemoteRequests *bool
	CacheSize           string
	Extra               map[string]string
}

type IPService struct {
	ID          string
	Name        string
	Port        string
	Address     string
	Certificate string
	TLSVersion  string
	Disabled    bool
	Invalid     bool
	Raw         map[string]string
}

type IPServiceSet struct {
	Port        string
	Address     string
	Certificate string
	TLSVersion  string
	Disabled    *bool
	Extra       map[string]string
}

type ARPEntry struct {
	ID         string
	Address    string
	MACAddress string
	Interface  string
	Complete   bool
	Dynamic    bool
	Disabled   bool
	Comment    string
	Raw        map[string]string
}

type ARPSet struct {
	Address    string
	MACAddress string
	Interface  string
	Comment    string
	Disabled   *bool
	Extra      map[string]string
}

type Pool struct {
	ID       string
	Name     string
	Ranges   string
	NextPool string
	Comment  string
	Raw      map[string]string
}

type PoolSet struct {
	Name     string
	Ranges   string
	NextPool string
	Comment  string
	Extra    map[string]string
}
