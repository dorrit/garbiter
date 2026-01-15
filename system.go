package garbiter

import (
	"strconv"

	"github.com/dorrit/garbiter/model"
)

type SystemAPI struct {
	svc model.Transport
}

// PrintIdentity returns the device identity in a typed struct.
func (s *SystemAPI) PrintIdentity() (*model.Identity, error) {
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
	_, err := s.svc.Run("/system/identity/set", "=name="+name)
	return err
}

// PrintResource returns the device resource information in a typed struct.
func (s *SystemAPI) PrintResource() (*model.Resource, error) {
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
