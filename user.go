package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type UserAPI struct {
	svc model.Transport
}

func (api *UserAPI) Print() ([]model.User, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/user/print", proplist(".id", "name", "group", "address", "last-logged-in", "disabled", "comment"))
	if err != nil {
		return nil, err
	}
	items := make([]model.User, 0, len(rows))
	for _, row := range rows {
		items = append(items, userFromMap(row))
	}
	return items, nil
}

func (api *UserAPI) Add(set model.UserSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("name", set.Name); err != nil {
		return nil, err
	}
	return api.svc.Run("/user/add", userSetArgs(set)...)
}

func (api *UserAPI) Set(id string, set model.UserSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/user/set", setIDArgs(id, userSetArgs(set))...)
	return err
}

func (api *UserAPI) Remove(id string) error {
	return api.runID("/user/remove", id)
}

func (api *UserAPI) Enable(id string) error {
	return api.runID("/user/enable", id)
}

func (api *UserAPI) Disable(id string) error {
	return api.runID("/user/disable", id)
}

func (api *UserAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func userFromMap(row map[string]string) model.User {
	return model.User{ID: row[".id"], Name: row["name"], Group: row["group"], Address: row["address"], LastLoggedIn: row["last-logged-in"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func userSetArgs(set model.UserSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "password", set.Password)
	args = appendArg(args, "group", set.Group)
	args = appendArg(args, "address", set.Address)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
