package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type IPAPI struct {
	svc model.Transport
}

func (api *IPAPI) Addresses() ([]model.IPAddress, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/address/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.IPAddress, 0, len(rows))
	for _, row := range rows {
		items = append(items, ipAddressFromMap(row))
	}
	return items, nil
}

func (api *IPAPI) AddAddress(set model.IPAddressSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ip/address/add", ipAddressSetArgs(set)...)
}

func (api *IPAPI) SetAddress(id string, set model.IPAddressSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/address/set", setIDArgs(id, ipAddressSetArgs(set))...)
	return err
}

func (api *IPAPI) RemoveAddress(id string) error {
	return api.runID("/ip/address/remove", id)
}

func (api *IPAPI) Routes() ([]model.Route, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/route/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.Route, 0, len(rows))
	for _, row := range rows {
		items = append(items, routeFromMap(row))
	}
	return items, nil
}

func (api *IPAPI) AddRoute(set model.RouteSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ip/route/add", routeSetArgs(set)...)
}

func (api *IPAPI) SetRoute(id string, set model.RouteSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/route/set", setIDArgs(id, routeSetArgs(set))...)
	return err
}

func (api *IPAPI) RemoveRoute(id string) error {
	return api.runID("/ip/route/remove", id)
}

func (api *IPAPI) PrintDNS() (*model.DNS, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	row, err := api.svc.Run("/ip/dns/print")
	if err != nil {
		return nil, err
	}
	dns := dnsFromMap(row)
	return &dns, nil
}

func (api *IPAPI) SetDNS(set model.DNSSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/dns/set", dnsSetArgs(set)...)
	return err
}

func (api *IPAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func ipAddressFromMap(row map[string]string) model.IPAddress {
	return model.IPAddress{ID: row[".id"], Address: row["address"], Network: row["network"], Interface: row["interface"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func routeFromMap(row map[string]string) model.Route {
	return model.Route{ID: row[".id"], Dst: row["dst-address"], Gateway: row["gateway"], Distance: row["distance"], Scope: row["scope"], Target: row["target-scope"], Active: boolFromRouterOS(row["active"]), Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func dnsFromMap(row map[string]string) model.DNS {
	return model.DNS{Servers: row["servers"], DynamicServers: row["dynamic-servers"], AllowRemoteRequests: boolFromRouterOS(row["allow-remote-requests"]), CacheSize: row["cache-size"], Raw: row}
}

func ipAddressSetArgs(set model.IPAddressSet) []string {
	args := []string{}
	args = appendArg(args, "address", set.Address)
	args = appendArg(args, "network", set.Network)
	args = appendArg(args, "interface", set.Interface)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func routeSetArgs(set model.RouteSet) []string {
	args := []string{}
	args = appendArg(args, "dst-address", set.Dst)
	args = appendArg(args, "gateway", set.Gateway)
	args = appendArg(args, "distance", set.Distance)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func dnsSetArgs(set model.DNSSet) []string {
	args := []string{}
	args = appendArg(args, "servers", set.Servers)
	args = appendArg(args, "cache-size", set.CacheSize)
	args = appendBoolArg(args, "allow-remote-requests", set.AllowRemoteRequests)
	return appendExtraArgs(args, set.Extra)
}
