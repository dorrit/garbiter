package garbiter

import (
	"fmt"
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

	cpuCount, err := parseOptionalInt("cpu-count", res["cpu-count"])
	if err != nil {
		return nil, err
	}
	writeSinceBoot, err := parseOptionalInt64("write-sect-since-reboot", res["write-sect-since-reboot"])
	if err != nil {
		return nil, err
	}
	writeTotal, err := parseOptionalInt64("write-sect-total", res["write-sect-total"])
	if err != nil {
		return nil, err
	}

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

	rows, err := s.svc.RunList("/system/health/print")
	if err != nil {
		return nil, err
	}

	health := &model.Health{Sensors: make([]model.HealthSensor, 0, len(rows))}
	for _, row := range rows {
		if name := row["name"]; name != "" {
			value, err := parseHealthNumber(name, row["value"])
			if err != nil {
				return nil, err
			}
			health.Sensors = append(health.Sensors, model.HealthSensor{
				Name:     name,
				Value:    value,
				RawValue: row["value"],
				Type:     row["type"],
				Raw:      row,
			})
			applyHealthValue(health, name, value)
			continue
		}

		for _, name := range []string{
			"voltage", "temperature", "power-consumption", "cpu-temperature",
			"fan1-speed", "fan2-speed", "fan3-speed", "fan4-speed",
			"board-temp1", "board-temp2", "board-temperature1", "board-temperature2",
			"psu1-voltage", "psu2-voltage", "psu1-current", "psu2-current",
		} {
			if row[name] == "" {
				continue
			}
			value, err := parseHealthNumber(name, row[name])
			if err != nil {
				return nil, err
			}
			health.Sensors = append(health.Sensors, model.HealthSensor{Name: name, Value: value, RawValue: row[name], Raw: row})
			applyHealthValue(health, name, value)
		}
	}

	return health, nil
}

func (s *SystemAPI) SetHealth(set model.HealthSettings) error {
	if s == nil || s.svc == nil {
		return service.ErrNotConnected
	}

	if set.CPUOvertempStartupDelay != nil && *set.CPUOvertempStartupDelay < 0 {
		return fmt.Errorf("cpu-overtemp-startup-delay must not be negative")
	}

	args := []string{}
	args = appendBoolArg(args, "cpu-overtemp-check", set.CPUOvertempCheck)
	args = appendIntArg(args, "cpu-overtemp-threshold", set.CPUOvertempThreshold)
	args = appendDurationArg(args, "cpu-overtemp-startup-delay", set.CPUOvertempStartupDelay)
	args = appendStringPtrArg(args, "fan-mode", set.FanMode)
	args = appendIntArg(args, "fan-on-threshold", set.FanOnThreshold)
	args = appendStringPtrArg(args, "fan-switch", set.FanSwitch)
	args = appendBoolArg(args, "use-fan", set.UseFan)
	args = appendIntArg(args, "fan-target-temp", set.FanTargetTemp)
	args = appendIntArg(args, "fan-full-speed-temp", set.FanFullSpeedTemp)
	args = appendIntArg(args, "fan-min-speed-percent", set.FanMinSpeedPercent)
	args = appendExtraArgs(args, set.Extra)
	_, err := s.svc.Run("/system/health/settings/set", args...)
	return err
}

func parseHealthNumber(name, raw string) (float64, error) {
	value, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, fmt.Errorf("parse health sensor %q value %q: %w", name, raw, err)
	}
	return value, nil
}

func parseOptionalInt(name, raw string) (int, error) {
	if raw == "" {
		return 0, nil
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parse %s value %q: %w", name, raw, err)
	}
	return value, nil
}

func parseOptionalInt64(name, raw string) (int64, error) {
	if raw == "" {
		return 0, nil
	}
	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse %s value %q: %w", name, raw, err)
	}
	return value, nil
}

func applyHealthValue(health *model.Health, name string, value float64) {
	switch name {
	case "voltage":
		health.Voltage = value
	case "temperature":
		health.Temperature = value
	case "power-consumption":
		health.PowerConsumption = value
	case "cpu-temperature":
		health.CPUTemperature = value
	case "fan1-speed":
		health.Fan1Speed = int(value)
	case "fan2-speed":
		health.Fan2Speed = int(value)
	case "fan3-speed":
		health.Fan3Speed = int(value)
	case "fan4-speed":
		health.Fan4Speed = int(value)
	case "board-temp1", "board-temperature1":
		health.BoardTemp1 = value
	case "board-temp2", "board-temperature2":
		health.BoardTemp2 = value
	case "psu1-voltage":
		health.PSU1Voltage = value
	case "psu2-voltage":
		health.PSU2Voltage = value
	case "psu1-current":
		health.PSU1Current = value
	case "psu2-current":
		health.PSU2Current = value
	}
}

func (s *SystemAPI) PrintClock() (*model.Clock, error) {
	if s == nil || s.svc == nil {
		return nil, service.ErrNotConnected
	}

	res, err := s.svc.Run("/system/clock/print")
	if err != nil {
		return nil, err
	}

	return &model.Clock{
		Time:         res["time"],
		Date:         res["date"],
		TimeZoneName: res["time-zone-name"],
		GMTOffset:    res["gmt-offset"],
		DSTActive:    boolFromRouterOS(res["dst-active"]),
		TimeZoneAuto: boolFromRouterOS(res["time-zone-autodetect"]),
		Raw:          res,
	}, nil
}

func (s *SystemAPI) SetClock(set model.ClockSet) error {
	if s == nil || s.svc == nil {
		return service.ErrNotConnected
	}

	args := []string{}
	args = appendArg(args, "time", set.Time)
	args = appendArg(args, "date", set.Date)
	args = appendArg(args, "time-zone-name", set.TimeZoneName)
	args = appendBoolArg(args, "time-zone-autodetect", set.TimeZoneAuto)
	args = appendExtraArgs(args, set.Extra)
	_, err := s.svc.Run("/system/clock/set", args...)
	return err
}

func (s *SystemAPI) Packages() ([]model.Package, error) {
	if s == nil || s.svc == nil {
		return nil, service.ErrNotConnected
	}

	rows, err := s.svc.RunList("/system/package/print")
	if err != nil {
		return nil, err
	}

	items := make([]model.Package, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.Package{
			ID:        row[".id"],
			Name:      row["name"],
			Version:   row["version"],
			BuildTime: row["build-time"],
			Scheduled: row["scheduled"],
			Disabled:  boolFromRouterOS(row["disabled"]),
			Raw:       row,
		})
	}
	return items, nil
}

func (s *SystemAPI) PrintRouterboard() (*model.Routerboard, error) {
	if s == nil || s.svc == nil {
		return nil, service.ErrNotConnected
	}

	res, err := s.svc.Run("/system/routerboard/print")
	if err != nil {
		return nil, err
	}

	currentFirmware := res["current-firmware"]
	upgradeFirmware := res["upgrade-firmware"]
	return &model.Routerboard{
		Routerboard:           boolFromRouterOS(res["routerboard"]),
		Model:                 res["model"],
		SerialNumber:          res["serial-number"],
		FirmwareType:          res["firmware-type"],
		FactoryFirmware:       res["factory-firmware"],
		CurrentFirmware:       currentFirmware,
		UpgradeFirmware:       upgradeFirmware,
		FirmwareUpgradeNeeded: currentFirmware != "" && upgradeFirmware != "" && currentFirmware != upgradeFirmware,
		Raw:                   res,
	}, nil
}

func (s *SystemAPI) SetRouterboardSettings(set model.RouterboardSettings) error {
	if s == nil || s.svc == nil {
		return service.ErrNotConnected
	}

	args := []string{}
	args = appendArg(args, "boot-device", set.BootDevice)
	args = appendArg(args, "cpu-frequency", set.CPUFrequency)
	args = appendArg(args, "boot-protocol", set.BootProtocol)
	args = appendEnabledArg(args, "protected-routerboot", set.ProtectedRouterboot)
	args = appendExtraArgs(args, set.Extra)
	_, err := s.svc.Run("/system/routerboard/settings/set", args...)
	return err
}

func (s *SystemAPI) UpgradeRouterboard() error {
	if s == nil || s.svc == nil {
		return service.ErrNotConnected
	}

	_, err := s.svc.Run("/system/routerboard/upgrade")
	return err
}
