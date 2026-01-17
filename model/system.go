package model

import "time"

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

type Health struct {
	// Common hardware monitoring information available on most devices
	Voltage     float64
	Temperature float64

	// Everything below is device specific and may not be populated
	// Example on CCR1072-1G-8S+ device:
	PowerConsumption float64
	CPUTemperature   float64

	Fan1Speed int
	Fan2Speed int
	Fan3Speed int
	Fan4Speed int

	BoardTemp1 float64
	BoardTemp2 float64

	PSU1Voltage float64
	PSU2Voltage float64

	PSU1Current float64
	PSU2Current float64
}

type HealthSettings struct {
	// CPU protection (common)
	CPUOvertempCheck        bool
	CPUOvertempThreshold    int
	CPUOvertempStartupDelay time.Duration

	// Fan (device-dependent)
	FanMode        string
	FanOnThreshold int
	FanSwitch      string
	UseFan         bool

	// Future-proof
	Extra map[string]string
}
