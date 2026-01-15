package garbiter

import "github.com/dorrit/garbiter/model"

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

	return &model.Resource{
		CPU:         res["cpu"],
		CPULoad:     res["cpu-load"],
		FreeMemory:  res["free-memory"],
		TotalMemory: res["total-memory"],
		Uptime:      res["uptime"],
	}, nil
}
