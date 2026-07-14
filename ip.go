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
	rows, err := api.svc.RunList("/ip/address/print", proplist(".id", "address", "network", "interface", "disabled", "comment"))
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
	if err := requireField("address", set.Address); err != nil {
		return nil, err
	}
	if err := requireField("interface", set.Interface); err != nil {
		return nil, err
	}
	return api.svc.Run("/ip/address/add", ipAddressSetArgs(set)...)
}

func (api *IPAPI) SetAddress(id string, set model.IPAddressSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
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
	rows, err := api.svc.RunList("/ip/route/print", proplist(".id", "dst-address", "gateway", "distance", "scope", "target-scope", "active", "disabled", "comment"))
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
	if err := validateID(id); err != nil {
		return err
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
	row, err := api.svc.Run("/ip/dns/print", proplist("servers", "dynamic-servers", "allow-remote-requests", "cache-size"))
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

func (api *IPAPI) Services() ([]model.IPService, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/service/print", proplist(".id", "name", "port", "address", "certificate", "tls-version", "disabled", "invalid"))
	if err != nil {
		return nil, err
	}
	items := make([]model.IPService, 0, len(rows))
	for _, row := range rows {
		items = append(items, ipServiceFromMap(row))
	}
	return items, nil
}

func (api *IPAPI) SetService(id string, set model.IPServiceSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/ip/service/set", setIDArgs(id, ipServiceSetArgs(set))...)
	return err
}

func (api *IPAPI) EnableService(id string) error {
	return api.runID("/ip/service/enable", id)
}

func (api *IPAPI) DisableService(id string) error {
	return api.runID("/ip/service/disable", id)
}

func (api *IPAPI) ARP() ([]model.ARPEntry, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/arp/print", proplist(".id", "address", "mac-address", "interface", "complete", "dynamic", "disabled", "comment"))
	if err != nil {
		return nil, err
	}
	items := make([]model.ARPEntry, 0, len(rows))
	for _, row := range rows {
		items = append(items, arpFromMap(row))
	}
	return items, nil
}

func (api *IPAPI) AddARP(set model.ARPSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("address", set.Address); err != nil {
		return nil, err
	}
	if err := requireField("mac-address", set.MACAddress); err != nil {
		return nil, err
	}
	if err := requireField("interface", set.Interface); err != nil {
		return nil, err
	}
	return api.svc.Run("/ip/arp/add", arpSetArgs(set)...)
}

func (api *IPAPI) SetARP(id string, set model.ARPSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/ip/arp/set", setIDArgs(id, arpSetArgs(set))...)
	return err
}

func (api *IPAPI) RemoveARP(id string) error {
	return api.runID("/ip/arp/remove", id)
}

func (api *IPAPI) Pools() ([]model.Pool, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/pool/print", proplist(".id", "name", "ranges", "next-pool", "comment"))
	if err != nil {
		return nil, err
	}
	items := make([]model.Pool, 0, len(rows))
	for _, row := range rows {
		items = append(items, poolFromMap(row))
	}
	return items, nil
}

func (api *IPAPI) AddPool(set model.PoolSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	if err := requireField("name", set.Name); err != nil {
		return nil, err
	}
	if err := requireField("ranges", set.Ranges); err != nil {
		return nil, err
	}
	return api.svc.Run("/ip/pool/add", poolSetArgs(set)...)
}

func (api *IPAPI) SetPool(id string, set model.PoolSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
	}
	_, err := api.svc.Run("/ip/pool/set", setIDArgs(id, poolSetArgs(set))...)
	return err
}

func (api *IPAPI) RemovePool(id string) error {
	return api.runID("/ip/pool/remove", id)
}

func (api *IPAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	if err := validateID(id); err != nil {
		return err
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

func ipServiceFromMap(row map[string]string) model.IPService {
	return model.IPService{ID: row[".id"], Name: row["name"], Port: row["port"], Address: row["address"], Certificate: row["certificate"], TLSVersion: row["tls-version"], Disabled: boolFromRouterOS(row["disabled"]), Invalid: boolFromRouterOS(row["invalid"]), Raw: row}
}

func arpFromMap(row map[string]string) model.ARPEntry {
	return model.ARPEntry{ID: row[".id"], Address: row["address"], MACAddress: row["mac-address"], Interface: row["interface"], Complete: boolFromRouterOS(row["complete"]), Dynamic: boolFromRouterOS(row["dynamic"]), Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func poolFromMap(row map[string]string) model.Pool {
	return model.Pool{ID: row[".id"], Name: row["name"], Ranges: row["ranges"], NextPool: row["next-pool"], Comment: row["comment"], Raw: row}
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

func ipServiceSetArgs(set model.IPServiceSet) []string {
	args := []string{}
	args = appendArg(args, "port", set.Port)
	args = appendArg(args, "address", set.Address)
	args = appendArg(args, "certificate", set.Certificate)
	args = appendArg(args, "tls-version", set.TLSVersion)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func arpSetArgs(set model.ARPSet) []string {
	args := []string{}
	args = appendArg(args, "address", set.Address)
	args = appendArg(args, "mac-address", set.MACAddress)
	args = appendArg(args, "interface", set.Interface)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func poolSetArgs(set model.PoolSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "ranges", set.Ranges)
	args = appendArg(args, "next-pool", set.NextPool)
	args = appendArg(args, "comment", set.Comment)
	return appendExtraArgs(args, set.Extra)
}
