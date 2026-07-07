package model

type FirewallRule struct {
	ID           string
	Chain        string
	Action       string
	Protocol     string
	SrcAddress   string
	DstAddress   string
	SrcPort      string
	DstPort      string
	InInterface  string
	OutInterface string
	Disabled     bool
	Dynamic      bool
	Comment      string
	Raw          map[string]string
}

type FirewallRuleSet struct {
	Chain        string
	Action       string
	Protocol     string
	SrcAddress   string
	DstAddress   string
	SrcPort      string
	DstPort      string
	InInterface  string
	OutInterface string
	Comment      string
	Disabled     *bool
	Extra        map[string]string
}

type AddressListEntry struct {
	ID       string
	List     string
	Address  string
	Timeout  string
	Dynamic  bool
	Disabled bool
	Comment  string
	Raw      map[string]string
}

type AddressListSet struct {
	List     string
	Address  string
	Timeout  string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}
