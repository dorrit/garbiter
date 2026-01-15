package model

type Identity struct {
	Name string
}

type Resource struct {
	Uptime          string
	Version         string
	BuildTime       string
	FactorySoftware string

	FreeMemory  string
	TotalMemory string

	CPU          string
	CPUCount     int
	CPUFrequency string
	CPULoad      string

	FreeHDDSpace  string
	TotalHDDSpace string

	WriteSectSinceBoot int64
	WriteSectTotal     int64

	BadBlocks string

	Architecture string
	BoardName    string
	Platform     string
}
