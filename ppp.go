package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type PPPAPI struct {
	svc model.Transport
}

func (api *PPPAPI) Profiles() ([]model.PPPProfile, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ppp/profile/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.PPPProfile, 0, len(rows))
	for _, row := range rows {
		items = append(items, pppProfileFromMap(row))
	}
	return items, nil
}

func (api *PPPAPI) AddProfile(set model.PPPProfileSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ppp/profile/add", pppProfileSetArgs(set)...)
}

func (api *PPPAPI) SetProfile(id string, set model.PPPProfileSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ppp/profile/set", setIDArgs(id, pppProfileSetArgs(set))...)
	return err
}

func (api *PPPAPI) RemoveProfile(id string) error {
	return api.runID("/ppp/profile/remove", id)
}

func (api *PPPAPI) Secrets() ([]model.PPPSecret, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ppp/secret/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.PPPSecret, 0, len(rows))
	for _, row := range rows {
		items = append(items, pppSecretFromMap(row))
	}
	return items, nil
}

func (api *PPPAPI) AddSecret(set model.PPPSecretSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ppp/secret/add", pppSecretSetArgs(set)...)
}

func (api *PPPAPI) SetSecret(id string, set model.PPPSecretSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ppp/secret/set", setIDArgs(id, pppSecretSetArgs(set))...)
	return err
}

func (api *PPPAPI) RemoveSecret(id string) error {
	return api.runID("/ppp/secret/remove", id)
}

func (api *PPPAPI) EnableSecret(id string) error {
	return api.runID("/ppp/secret/enable", id)
}

func (api *PPPAPI) DisableSecret(id string) error {
	return api.runID("/ppp/secret/disable", id)
}

func (api *PPPAPI) Active() ([]model.PPPActive, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ppp/active/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.PPPActive, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.PPPActive{ID: row[".id"], Name: row["name"], Service: row["service"], CallerID: row["caller-id"], Address: row["address"], Uptime: row["uptime"], Encoding: row["encoding"], Raw: row})
	}
	return items, nil
}

func (api *PPPAPI) RemoveActive(id string) error {
	return api.runID("/ppp/active/remove", id)
}

func (api *PPPAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func pppProfileFromMap(row map[string]string) model.PPPProfile {
	return model.PPPProfile{ID: row[".id"], Name: row["name"], LocalAddress: row["local-address"], RemoteAddress: row["remote-address"], DNSAddress: row["dns-server"], OnlyOne: row["only-one"], Comment: row["comment"], Raw: row}
}

func pppSecretFromMap(row map[string]string) model.PPPSecret {
	return model.PPPSecret{ID: row[".id"], Name: row["name"], Service: row["service"], Profile: row["profile"], RemoteAddress: row["remote-address"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func pppProfileSetArgs(set model.PPPProfileSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "local-address", set.LocalAddress)
	args = appendArg(args, "remote-address", set.RemoteAddress)
	args = appendArg(args, "dns-server", set.DNSAddress)
	args = appendArg(args, "only-one", set.OnlyOne)
	args = appendArg(args, "comment", set.Comment)
	return appendExtraArgs(args, set.Extra)
}

func pppSecretSetArgs(set model.PPPSecretSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "password", set.Password)
	args = appendArg(args, "service", set.Service)
	args = appendArg(args, "profile", set.Profile)
	args = appendArg(args, "remote-address", set.RemoteAddress)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
