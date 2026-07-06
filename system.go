package garbiter

import (
	"strconv"

	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type SystemAPI struct {
	svc model.Transport
}

// PrintIdentity returns the device identity in a typed struct.
func (s *SystemAPI) PrintIdentity() (*model.Identity, error) {
	if s == nil || s.svc == nil {
		return nil, service.ErrNotConnected
	}

	res, err := s.svc.Run("/system/identity/print")
	if err != nil {
		return nil, err
	}

	return &model.Identity{
		Name: res["name"],
	}, nil
}

// SetIdentity sets the device identity.
func (s *SystemAPI) SetIdentity(name string) error {
	if s == nil || s.svc == nil {
		return service.ErrNotConnected
	}

	_, err := s.svc.Run("/system/identity/set", "=name="+name)
	return err
}

// PrintResource returns the device resource information in a typed struct.
func (s *SystemAPI) PrintResource() (*model.Resource, error) {
	if s == nil || s.svc == nil {
		return nil, service.ErrNotConnected
	}

	res, err := s.svc.Run("/system/resource/print")
	if err != nil {
		return nil, err
	}

	cpuCount, _ := strconv.Atoi(res["cpu-count"])
	writeSinceBoot, _ := strconv.ParseInt(res["write-sect-since-reboot"], 10, 64)
	writeTotal, _ := strconv.ParseInt(res["write-sect-total"], 10, 64)

	return &model.Resource{
		Uptime: res["uptime"],

		Version:         res["version"],
		BuildTime:       res["build-time"],
		FactorySoftware: res["factory-software"],

		FreeMemory:  res["free-memory"],
		TotalMemory: res["total-memory"],

		CPU:          res["cpu"],
		CPUCount:     cpuCount,
		CPUFrequency: res["cpu-frequency"],
		CPULoad:      res["cpu-load"],

		FreeHDDSpace:  res["free-hdd-space"],
		TotalHDDSpace: res["total-hdd-space"],

		WriteSectSinceBoot: writeSinceBoot,
		WriteSectTotal:     writeTotal,

		BadBlocks: res["bad-blocks"],

		Architecture: res["architecture-name"],
		BoardName:    res["board-name"],
		Platform:     res["platform"],
	}, nil
}

// PrintHealth returns the device health information in a typed struct.
func (s *SystemAPI) PrintHealth() (*model.Health, error) {
	if s == nil || s.svc == nil {
		return nil, service.ErrNotConnected
	}

	res, err := s.svc.Run("/system/health/print")
	if err != nil {
		return nil, err
	}
	powerConsumption, _ := strconv.ParseFloat(res["power-consumption"], 64)
	cpuTemperature, _ := strconv.ParseFloat(res["cpu-temperature"], 64)
	fan1Speed, _ := strconv.Atoi(res["fan1-speed"])
	fan2Speed, _ := strconv.Atoi(res["fan2-speed"])
	fan3Speed, _ := strconv.Atoi(res["fan3-speed"])
	fan4Speed, _ := strconv.Atoi(res["fan4-speed"])
	boardTemp1, _ := strconv.ParseFloat(res["board-temp1"], 64)
	boardTemp2, _ := strconv.ParseFloat(res["board-temp2"], 64)
	psu1Voltage, _ := strconv.ParseFloat(res["psu1-voltage"], 64)
	psu2Voltage, _ := strconv.ParseFloat(res["psu2-voltage"], 64)
	psu1Current, _ := strconv.ParseFloat(res["psu1-current"], 64)
	psu2Current, _ := strconv.ParseFloat(res["psu2-current"], 64)

	voltage, _ := strconv.ParseFloat(res["voltage"], 64)
	temperature, _ := strconv.ParseFloat(res["temperature"], 64)
	return &model.Health{
		Voltage:          voltage,
		Temperature:      temperature,
		PowerConsumption: powerConsumption,
		CPUTemperature:   cpuTemperature,
		Fan1Speed:        fan1Speed,
		Fan2Speed:        fan2Speed,
		Fan3Speed:        fan3Speed,
		Fan4Speed:        fan4Speed,
		BoardTemp1:       boardTemp1,
		BoardTemp2:       boardTemp2,
		PSU1Voltage:      psu1Voltage,
		PSU2Voltage:      psu2Voltage,
		PSU1Current:      psu1Current,
		PSU2Current:      psu2Current,
	}, nil
}

func (s *SystemAPI) SetHealth(set model.HealthSettings) error {
	if s == nil || s.svc == nil {
		return service.ErrNotConnected
	}

	args := []string{}
	if set.CPUOvertempCheck {
		args = append(args, "=cpu-overtemp-check=yes")
	}
	if set.CPUOvertempThreshold > 0 {
		args = append(args, "=cpu-overtemp-threshold="+strconv.Itoa(set.CPUOvertempThreshold))
	}
	if set.CPUOvertempStartupDelay.Abs() > 0 {
		args = append(args, "=cpu-overtemp-startup-delay="+set.CPUOvertempStartupDelay.String())
	}
	if set.FanMode != "" {
		args = append(args, "=fan-mode="+set.FanMode)
	}
	if FanOnThreshold := set.FanOnThreshold; FanOnThreshold > 0 {
		args = append(args, "=fan-on-threshold="+strconv.Itoa(FanOnThreshold))
	}
	if set.FanSwitch != "" {
		args = append(args, "=fan-switch="+set.FanSwitch)
	}
	if set.UseFan {
		args = append(args, "=use-fan=yes")
	} else {
		args = append(args, "=use-fan=no")
	}

	if set.Extra != nil {
		for k, v := range set.Extra {
			args = append(args, "="+k+"="+v)
		}
	}

	_, err := s.svc.Run("/system/health/settings/set", args...)
	return err
}
