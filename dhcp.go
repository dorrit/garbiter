package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type DHCPAPI struct {
	svc model.Transport
}

func (api *DHCPAPI) Clients() ([]model.DHCPClient, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/dhcp-client/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.DHCPClient, 0, len(rows))
	for _, row := range rows {
		items = append(items, dhcpClientFromMap(row))
	}
	return items, nil
}

func (api *DHCPAPI) AddClient(set model.DHCPClientSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ip/dhcp-client/add", dhcpClientSetArgs(set)...)
}

func (api *DHCPAPI) SetClient(id string, set model.DHCPClientSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/dhcp-client/set", setIDArgs(id, dhcpClientSetArgs(set))...)
	return err
}

func (api *DHCPAPI) RemoveClient(id string) error {
	return api.runID("/ip/dhcp-client/remove", id)
}

func (api *DHCPAPI) Servers() ([]model.DHCPServer, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/dhcp-server/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.DHCPServer, 0, len(rows))
	for _, row := range rows {
		items = append(items, dhcpServerFromMap(row))
	}
	return items, nil
}

func (api *DHCPAPI) AddServer(set model.DHCPServerSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/ip/dhcp-server/add", dhcpServerSetArgs(set)...)
}

func (api *DHCPAPI) SetServer(id string, set model.DHCPServerSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/ip/dhcp-server/set", setIDArgs(id, dhcpServerSetArgs(set))...)
	return err
}

func (api *DHCPAPI) RemoveServer(id string) error {
	return api.runID("/ip/dhcp-server/remove", id)
}

func (api *DHCPAPI) Leases() ([]model.DHCPLease, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ip/dhcp-server/lease/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.DHCPLease, 0, len(rows))
	for _, row := range rows {
		items = append(items, dhcpLeaseFromMap(row))
	}
	return items, nil
}

func (api *DHCPAPI) MakeLeaseStatic(id string) error {
	return api.runID("/ip/dhcp-server/lease/make-static", id)
}

func (api *DHCPAPI) RemoveLease(id string) error {
	return api.runID("/ip/dhcp-server/lease/remove", id)
}

func (api *DHCPAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func dhcpClientFromMap(row map[string]string) model.DHCPClient {
	return model.DHCPClient{ID: row[".id"], Interface: row["interface"], Address: row["address"], Gateway: row["gateway"], Status: row["status"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func dhcpServerFromMap(row map[string]string) model.DHCPServer {
	return model.DHCPServer{ID: row[".id"], Name: row["name"], Interface: row["interface"], AddressPool: row["address-pool"], LeaseTime: row["lease-time"], Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func dhcpLeaseFromMap(row map[string]string) model.DHCPLease {
	return model.DHCPLease{ID: row[".id"], Address: row["address"], MACAddress: row["mac-address"], HostName: row["host-name"], Server: row["server"], Status: row["status"], Dynamic: boolFromRouterOS(row["dynamic"]), Disabled: boolFromRouterOS(row["disabled"]), Comment: row["comment"], Raw: row}
}

func dhcpClientSetArgs(set model.DHCPClientSet) []string {
	args := []string{}
	args = appendArg(args, "interface", set.Interface)
	args = appendBoolArg(args, "use-peer-dns", set.UsePeerDNS)
	args = appendBoolArg(args, "use-peer-ntp", set.UsePeerNTP)
	args = appendArg(args, "add-default-route", set.AddDefaultRoute)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}

func dhcpServerSetArgs(set model.DHCPServerSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "interface", set.Interface)
	args = appendArg(args, "address-pool", set.AddressPool)
	args = appendArg(args, "lease-time", set.LeaseTime)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
