package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type ToolAPI struct {
	svc model.Transport
}

func (api *ToolAPI) Ping(req model.PingRequest) ([]model.PingResult, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/ping", pingArgs(req)...)
	if err != nil {
		return nil, err
	}
	items := make([]model.PingResult, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.PingResult{Host: row["host"], Seq: row["seq"], Size: row["size"], TTL: row["ttl"], Time: row["time"], Status: row["status"], Raw: row})
	}
	return items, nil
}

func pingArgs(req model.PingRequest) []string {
	args := []string{}
	args = appendArg(args, "address", req.Address)
	args = appendArg(args, "count", req.Count)
	args = appendArg(args, "interface", req.Interface)
	args = appendArg(args, "routing-table", req.RoutingTable)
	return appendExtraArgs(args, req.Extra)
}
