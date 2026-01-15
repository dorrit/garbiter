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
