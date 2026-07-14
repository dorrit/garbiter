package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type CertificateAPI struct {
	svc model.Transport
}

func (api *CertificateAPI) Print() ([]model.Certificate, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/certificate/print", proplist(".id", "name", "common-name", "fingerprint", "invalid-before", "invalid-after", "trusted", "private-key"))
	if err != nil {
		return nil, err
	}
	items := make([]model.Certificate, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.Certificate{ID: row[".id"], Name: row["name"], CommonName: row["common-name"], Fingerprint: row["fingerprint"], InvalidBefore: row["invalid-before"], InvalidAfter: row["invalid-after"], Trusted: boolFromRouterOS(row["trusted"]), PrivateKey: boolFromRouterOS(row["private-key"]), Raw: row})
	}
	return items, nil
}

func (api *CertificateAPI) Set(id string, set model.CertificateSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendBoolArg(args, "trusted", set.Trusted)
	args = appendExtraArgs(args, set.Extra)
	_, err := api.svc.Run("/certificate/set", setIDArgs(id, args)...)
	return err
}

func (api *CertificateAPI) Remove(id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/certificate/remove", "=.id="+id)
	return err
}

type SNMPAPI struct {
	svc model.Transport
}

func (api *SNMPAPI) Print() (*model.SNMP, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	row, err := api.svc.Run("/snmp/print", proplist("enabled", "contact", "location", "trap-community", "trap-version"))
	if err != nil {
		return nil, err
	}
	return &model.SNMP{Enabled: boolFromRouterOS(row["enabled"]), Contact: row["contact"], Location: row["location"], TrapCommunity: row["trap-community"], TrapVersion: row["trap-version"], Raw: row}, nil
}

func (api *SNMPAPI) Set(set model.SNMPSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	args := []string{}
	args = appendBoolArg(args, "enabled", set.Enabled)
	args = appendArg(args, "contact", set.Contact)
	args = appendArg(args, "location", set.Location)
	args = appendArg(args, "trap-community", set.TrapCommunity)
	args = appendArg(args, "trap-version", set.TrapVersion)
	args = appendExtraArgs(args, set.Extra)
	_, err := api.svc.Run("/snmp/set", args...)
	return err
}

type ScheduleAPI struct {
	svc model.Transport
}

func (api *ScheduleAPI) Print() ([]model.Schedule, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/system/scheduler/print", proplist(".id", "name", "start-date", "start-time", "interval", "on-event", "disabled", "comment"))
	if err != nil {
		return nil, err
	}
	items := make([]model.Schedule, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.Schedule{ID: row[".id"], Name: row["name"], StartDate: row["start-date"], StartTime: row["start-time"], Interval: row["interval"], OnEvent: row["on-event"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row})
	}
	return items, nil
}

func (api *ScheduleAPI) Add(set model.ScheduleSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("name", set.Name); err != nil {
		return nil, err
	}
	return api.svc.Run("/system/scheduler/add", scheduleSetArgs(set)...)
}

func (api *ScheduleAPI) Set(id string, set model.ScheduleSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/system/scheduler/set", setIDArgs(id, scheduleSetArgs(set))...)
	return err
}

func (api *ScheduleAPI) Remove(id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/system/scheduler/remove", "=.id="+id)
	return err
}

type ScriptAPI struct {
	svc model.Transport
}

func (api *ScriptAPI) Print() ([]model.Script, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/system/script/print", proplist(".id", "name", "owner", "policy", "run-count", "last-started", "source", "invalid"))
	if err != nil {
		return nil, err
	}
	items := make([]model.Script, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.Script{ID: row[".id"], Name: row["name"], Owner: row["owner"], Policy: row["policy"], RunCount: row["run-count"], LastStarted: row["last-started"], Source: row["source"], Invalid: boolFromRouterOS(row["invalid"]), Raw: row})
	}
	return items, nil
}

func (api *ScriptAPI) Add(set model.ScriptSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("name", set.Name); err != nil {
		return nil, err
	}
	return api.svc.Run("/system/script/add", scriptSetArgs(set)...)
}

func (api *ScriptAPI) Set(id string, set model.ScriptSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/system/script/set", setIDArgs(id, scriptSetArgs(set))...)
	return err
}

func (api *ScriptAPI) Remove(id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/system/script/remove", "=.id="+id)
	return err
}

func (api *ScriptAPI) Run(id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/system/script/run", "=.id="+id)
	return err
}

func scheduleSetArgs(set model.ScheduleSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "start-date", set.StartDate)
	args = appendArg(args, "start-time", set.StartTime)
	args = appendArg(args, "interval", set.Interval)
	args = appendArg(args, "on-event", set.OnEvent)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func scriptSetArgs(set model.ScriptSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "source", set.Source)
	args = appendArg(args, "policy", set.Policy)
	args = appendArg(args, "owner", set.Owner)
	return appendExtraArgs(args, set.Extra)
}
