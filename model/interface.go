package model

type Interface struct {
	ID         string
	Name       string
	Type       string
	ActualMTU  string
	L2MTU      string
	MACAddress string
	Running    bool
	Disabled   bool
	Comment    string
	Raw        map[string]string
}

type InterfaceSet struct {
	Name     string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}

type Bridge struct {
	ID       string
	Name     string
	Protocol string
	Disabled bool
	Comment  string
	Raw      map[string]string
}

type BridgeSet struct {
	Name     string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}

type VLAN struct {
	ID        string
	Name      string
	Interface string
	VLANID    string
	Disabled  bool
	Comment   string
	Raw       map[string]string
}

type VLANSet struct {
	Name      string
	Interface string
	VLANID    string
	Comment   string
	Disabled  *bool
	Extra     map[string]string
}
