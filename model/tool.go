package model

type PingRequest struct {
	Address      string
	Count        string
	Interface    string
	RoutingTable string
	Extra        map[string]string
}

type PingResult struct {
	Host   string
	Seq    string
	Size   string
	TTL    string
	Time   string
	Status string
	Raw    map[string]string
}
