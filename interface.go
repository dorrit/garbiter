package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type InterfaceAPI struct {
	svc model.Transport
}

func (api *InterfaceAPI) Print() ([]model.Interface, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/interface/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.Interface, 0, len(rows))
	for _, row := range rows {
		items = append(items, interfaceFromMap(row))
	}
	return items, nil
}

func (api *InterfaceAPI) Set(id string, set model.InterfaceSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/interface/set", setIDArgs(id, interfaceSetArgs(set))...)
	return err
}

func (api *InterfaceAPI) Enable(id string) error {
	return api.runID("/interface/enable", id)
}

func (api *InterfaceAPI) Disable(id string) error {
	return api.runID("/interface/disable", id)
}

func (api *InterfaceAPI) Bridges() ([]model.Bridge, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/interface/bridge/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.Bridge, 0, len(rows))
	for _, row := range rows {
		items = append(items, bridgeFromMap(row))
	}
	return items, nil
}

func (api *InterfaceAPI) AddBridge(set model.BridgeSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/interface/bridge/add", bridgeSetArgs(set)...)
}

func (api *InterfaceAPI) SetBridge(id string, set model.BridgeSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/interface/bridge/set", setIDArgs(id, bridgeSetArgs(set))...)
	return err
}

func (api *InterfaceAPI) RemoveBridge(id string) error {
	return api.runID("/interface/bridge/remove", id)
}

func (api *InterfaceAPI) VLANs() ([]model.VLAN, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/interface/vlan/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.VLAN, 0, len(rows))
	for _, row := range rows {
		items = append(items, vlanFromMap(row))
	}
	return items, nil
}

func (api *InterfaceAPI) AddVLAN(set model.VLANSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/interface/vlan/add", vlanSetArgs(set)...)
}

func (api *InterfaceAPI) SetVLAN(id string, set model.VLANSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/interface/vlan/set", setIDArgs(id, vlanSetArgs(set))...)
	return err
}

func (api *InterfaceAPI) RemoveVLAN(id string) error {
	return api.runID("/interface/vlan/remove", id)
}

func (api *InterfaceAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func interfaceFromMap(row map[string]string) model.Interface {
	return model.Interface{ID: row[".id"], Name: row["name"], Type: row["type"], ActualMTU: row["actual-mtu"], L2MTU: row["l2mtu"], MACAddress: row["mac-address"], Running: boolFromRouterOS(row["running"]), Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func bridgeFromMap(row map[string]string) model.Bridge {
	return model.Bridge{ID: row[".id"], Name: row["name"], Protocol: row["protocol-mode"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func vlanFromMap(row map[string]string) model.VLAN {
	return model.VLAN{ID: row[".id"], Name: row["name"], Interface: row["interface"], VLANID: row["vlan-id"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func interfaceSetArgs(set model.InterfaceSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func bridgeSetArgs(set model.BridgeSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func vlanSetArgs(set model.VLANSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "interface", set.Interface)
	args = appendArg(args, "vlan-id", set.VLANID)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
