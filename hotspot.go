package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type HotspotAPI struct {
	svc model.Transport
}

func (api *HotspotAPI) Servers() ([]model.HotspotServer, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/hotspot/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.HotspotServer, 0, len(rows))
	for _, row := range rows {
		items = append(items, hotspotServerFromMap(row))
	}
	return items, nil
}

func (api *HotspotAPI) AddServer(set model.HotspotServerSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ip/hotspot/add", hotspotServerSetArgs(set)...)
}

func (api *HotspotAPI) SetServer(id string, set model.HotspotServerSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/hotspot/set", setIDArgs(id, hotspotServerSetArgs(set))...)
	return err
}

func (api *HotspotAPI) RemoveServer(id string) error {
	return api.runID("/ip/hotspot/remove", id)
}

func (api *HotspotAPI) Users() ([]model.HotspotUser, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/hotspot/user/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.HotspotUser, 0, len(rows))
	for _, row := range rows {
		items = append(items, hotspotUserFromMap(row))
	}
	return items, nil
}

func (api *HotspotAPI) AddUser(set model.HotspotUserSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ip/hotspot/user/add", hotspotUserSetArgs(set)...)
}

func (api *HotspotAPI) SetUser(id string, set model.HotspotUserSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/hotspot/user/set", setIDArgs(id, hotspotUserSetArgs(set))...)
	return err
}

func (api *HotspotAPI) RemoveUser(id string) error {
	return api.runID("/ip/hotspot/user/remove", id)
}

func (api *HotspotAPI) EnableUser(id string) error {
	return api.runID("/ip/hotspot/user/enable", id)
}

func (api *HotspotAPI) DisableUser(id string) error {
	return api.runID("/ip/hotspot/user/disable", id)
}

func (api *HotspotAPI) Active() ([]model.HotspotActive, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/hotspot/active/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.HotspotActive, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.HotspotActive{ID: row[".id"], User: row["user"], Address: row["address"], MACAddress: row["mac-address"], Uptime: row["uptime"], Server: row["server"], Raw: row})
	}
	return items, nil
}

func (api *HotspotAPI) RemoveActive(id string) error {
	return api.runID("/ip/hotspot/active/remove", id)
}

func (api *HotspotAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func hotspotServerFromMap(row map[string]string) model.HotspotServer {
	return model.HotspotServer{ID: row[".id"], Name: row["name"], Interface: row["interface"], AddressPool: row["address-pool"], Profile: row["profile"], IdleTimeout: row["idle-timeout"], Disabled: boolFromRouterOS(row["disabled"]), Raw: row}
}

func hotspotUserFromMap(row map[string]string) model.HotspotUser {
	return model.HotspotUser{ID: row[".id"], Name: row["name"], Profile: row["profile"], Server: row["server"], Address: row["address"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func hotspotServerSetArgs(set model.HotspotServerSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "interface", set.Interface)
	args = appendArg(args, "address-pool", set.AddressPool)
	args = appendArg(args, "profile", set.Profile)
	args = appendArg(args, "idle-timeout", set.IdleTimeout)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func hotspotUserSetArgs(set model.HotspotUserSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "password", set.Password)
	args = appendArg(args, "profile", set.Profile)
	args = appendArg(args, "server", set.Server)
	args = appendArg(args, "address", set.Address)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
