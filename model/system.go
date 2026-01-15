package model

type Identity struct {
	Name string
}

type Resource struct {
	CPU         string
	CPULoad     string
	FreeMemory  string
	TotalMemory string
	Uptime      string
}
