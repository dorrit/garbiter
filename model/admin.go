package model

type Certificate struct {
	ID            string
	Name          string
	CommonName    string
	Fingerprint   string
	InvalidBefore string
	InvalidAfter  string
	Trusted       bool
	PrivateKey    bool
	Raw           map[string]string
}

type CertificateSet struct {
	Name    string
	Trusted *bool
	Extra   map[string]string
}

type SNMP struct {
	Enabled       bool
	Contact       string
	Location      string
	TrapCommunity string
	TrapVersion   string
	Raw           map[string]string
}

type SNMPSet struct {
	Enabled       *bool
	Contact       string
	Location      string
	TrapCommunity string
	TrapVersion   string
	Extra         map[string]string
}

type Schedule struct {
	ID        string
	Name      string
	StartDate string
	StartTime string
	Interval  string
	OnEvent   string
	Disabled  bool
	Comment   string
	Raw       map[string]string
}

type ScheduleSet struct {
	Name      string
	StartDate string
	StartTime string
	Interval  string
	OnEvent   string
	Comment   string
	Disabled  *bool
	Extra     map[string]string
}

type Script struct {
	ID          string
	Name        string
	Owner       string
	Policy      string
	RunCount    string
	LastStarted string
	Source      string
	Invalid     bool
	Raw         map[string]string
}

type ScriptSet struct {
	Name   string
	Source string
	Policy string
	Owner  string
	Extra  map[string]string
}
