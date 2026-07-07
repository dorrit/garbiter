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
