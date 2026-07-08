package model

type User struct {
	ID           string
	Name         string
	Group        string
	Address      string
	LastLoggedIn string
	Disabled     bool
	Comment      string
	Raw          map[string]string
}

type UserSet struct {
	Name     string
	Password string
	Group    string
	Address  string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}
