package garbiter

import (
	"strconv"

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
	if err := requireField("address", req.Address); err != nil {
		return nil, err
	}
	if req.Count != "" {
		count, err := strconv.Atoi(req.Count)
		if err != nil || count < 1 {
			return nil, &ValidationError{Field: "count", Message: "must be a positive integer"}
		}
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
	count := req.Count
	if count == "" {
		count = "4"
	}
	args = appendArg(args, "count", count)
	args = appendArg(args, "interface", req.Interface)
	args = appendArg(args, "routing-table", req.RoutingTable)
	return appendExtraArgs(args, req.Extra)
}
