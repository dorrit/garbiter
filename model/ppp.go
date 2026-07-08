package model

type PPPProfile struct {
	ID            string
	Name          string
	LocalAddress  string
	RemoteAddress string
	DNSAddress    string
	OnlyOne       string
	Comment       string
	Raw           map[string]string
}

type PPPProfileSet struct {
	Name          string
	LocalAddress  string
	RemoteAddress string
	DNSAddress    string
	OnlyOne       string
	Comment       string
	Extra         map[string]string
}

type PPPSecret struct {
	ID            string
	Name          string
	Service       string
	Profile       string
	RemoteAddress string
	Disabled      bool
	Comment       string
	Raw           map[string]string
}

type PPPSecretSet struct {
	Name          string
	Password      string
	Service       string
	Profile       string
	RemoteAddress string
	Comment       string
	Disabled      *bool
	Extra         map[string]string
}

type PPPActive struct {
	ID       string
	Name     string
	Service  string
	CallerID string
	Address  string
	Uptime   string
	Encoding string
	Raw      map[string]string
}
