package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type FirewallAPI struct {
	svc model.Transport
}

func (api *FirewallAPI) FilterRules() ([]model.FirewallRule, error) {
	return api.rules("/ip/firewall/filter/print")
}

func (api *FirewallAPI) AddFilterRule(set model.FirewallRuleSet) (map[string]string, error) {
	return api.addRule("/ip/firewall/filter/add", set)
}

func (api *FirewallAPI) SetFilterRule(id string, set model.FirewallRuleSet) error {
	return api.setRule("/ip/firewall/filter/set", id, set)
}

func (api *FirewallAPI) RemoveFilterRule(id string) error {
	return api.runID("/ip/firewall/filter/remove", id)
}

func (api *FirewallAPI) EnableFilterRule(id string) error {
	return api.runID("/ip/firewall/filter/enable", id)
}

func (api *FirewallAPI) DisableFilterRule(id string) error {
	return api.runID("/ip/firewall/filter/disable", id)
}

func (api *FirewallAPI) NATRules() ([]model.FirewallRule, error) {
	return api.rules("/ip/firewall/nat/print")
}

func (api *FirewallAPI) AddNATRule(set model.FirewallRuleSet) (map[string]string, error) {
	return api.addRule("/ip/firewall/nat/add", set)
}

func (api *FirewallAPI) SetNATRule(id string, set model.FirewallRuleSet) error {
	return api.setRule("/ip/firewall/nat/set", id, set)
}

func (api *FirewallAPI) RemoveNATRule(id string) error {
	return api.runID("/ip/firewall/nat/remove", id)
}

func (api *FirewallAPI) EnableNATRule(id string) error {
	return api.runID("/ip/firewall/nat/enable", id)
}

func (api *FirewallAPI) DisableNATRule(id string) error {
	return api.runID("/ip/firewall/nat/disable", id)
}

func (api *FirewallAPI) AddressList() ([]model.AddressListEntry, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/firewall/address-list/print", proplist(".id", "list", "address", "timeout", "dynamic", "disabled", "comment"))
	if err != nil {
		return nil, err
	}
	items := make([]model.AddressListEntry, 0, len(rows))
	for _, row := range rows {
		items = append(items, addressListFromMap(row))
	}
	return items, nil
}

func (api *FirewallAPI) AddAddressList(set model.AddressListSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("list", set.List); err != nil {
		return nil, err
	}
	if err := requireField("address", set.Address); err != nil {
		return nil, err
	}
	return api.svc.Run("/ip/firewall/address-list/add", addressListSetArgs(set)...)
}

func (api *FirewallAPI) SetAddressList(id string, set model.AddressListSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/ip/firewall/address-list/set", setIDArgs(id, addressListSetArgs(set))...)
	return err
}

func (api *FirewallAPI) RemoveAddressList(id string) error {
	return api.runID("/ip/firewall/address-list/remove", id)
}

func (api *FirewallAPI) rules(cmd string) ([]model.FirewallRule, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList(cmd, proplist(".id", "chain", "action", "protocol", "src-address", "dst-address", "src-port", "dst-port", "in-interface", "out-interface", "disabled", "dynamic", "comment"))
	if err != nil {
		return nil, err
	}
	items := make([]model.FirewallRule, 0, len(rows))
	for _, row := range rows {
		items = append(items, firewallRuleFromMap(row))
	}
	return items, nil
}

func (api *FirewallAPI) addRule(cmd string, set model.FirewallRuleSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("chain", set.Chain); err != nil {
		return nil, err
	}
	return api.svc.Run(cmd, firewallRuleSetArgs(set)...)
}

func (api *FirewallAPI) setRule(cmd, id string, set model.FirewallRuleSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run(cmd, setIDArgs(id, firewallRuleSetArgs(set))...)
	return err
}

func (api *FirewallAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func firewallRuleFromMap(row map[string]string) model.FirewallRule {
	return model.FirewallRule{ID: row[".id"], Chain: row["chain"], Action: row["action"], Protocol: row["protocol"], SrcAddress: row["src-address"], DstAddress: row["dst-address"], SrcPort: row["src-port"], DstPort: row["dst-port"], InInterface: row["in-interface"], OutInterface: row["out-interface"], Disabled: boolFromRouterOS(row["disabled"]), Dynamic: boolFromRouterOS(row["dynamic"]), Comment: row["comment"], Raw: row}
}

func addressListFromMap(row map[string]string) model.AddressListEntry {
	return model.AddressListEntry{ID: row[".id"], List: row["list"], Address: row["address"], Timeout: row["timeout"], Dynamic: boolFromRouterOS(row["dynamic"]), Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func firewallRuleSetArgs(set model.FirewallRuleSet) []string {
	args := []string{}
	args = appendArg(args, "chain", set.Chain)
	args = appendArg(args, "action", set.Action)
	args = appendArg(args, "protocol", set.Protocol)
	args = appendArg(args, "src-address", set.SrcAddress)
	args = appendArg(args, "dst-address", set.DstAddress)
	args = appendArg(args, "src-port", set.SrcPort)
	args = appendArg(args, "dst-port", set.DstPort)
	args = appendArg(args, "in-interface", set.InInterface)
	args = appendArg(args, "out-interface", set.OutInterface)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func addressListSetArgs(set model.AddressListSet) []string {
	args := []string{}
	args = appendArg(args, "list", set.List)
	args = appendArg(args, "address", set.Address)
	args = appendArg(args, "timeout", set.Timeout)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
