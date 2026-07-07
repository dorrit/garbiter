package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type QueueAPI struct {
	svc model.Transport
}

func (api *QueueAPI) Simple() ([]model.SimpleQueue, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/queue/simple/print")
	if err != nil {
		return nil, err
	}
	items := make([]model.SimpleQueue, 0, len(rows))
	for _, row := range rows {
		items = append(items, simpleQueueFromMap(row))
	}
	return items, nil
}

func (api *QueueAPI) AddSimple(set model.SimpleQueueSet) (map[string]string, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	return api.svc.Run("/queue/simple/add", simpleQueueSetArgs(set)...)
}

func (api *QueueAPI) SetSimple(id string, set model.SimpleQueueSet) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run("/queue/simple/set", setIDArgs(id, simpleQueueSetArgs(set))...)
	return err
}

func (api *QueueAPI) RemoveSimple(id string) error {
	return api.runID("/queue/simple/remove", id)
}

func (api *QueueAPI) EnableSimple(id string) error {
	return api.runID("/queue/simple/enable", id)
}

func (api *QueueAPI) DisableSimple(id string) error {
	return api.runID("/queue/simple/disable", id)
}

func (api *QueueAPI) runID(cmd, id string) error {
	if api == nil || api.svc == nil {
		return service.ErrNotConnected
	}
	_, err := api.svc.Run(cmd, "=.id="+id)
	return err
}

func simpleQueueFromMap(row map[string]string) model.SimpleQueue {
	return model.SimpleQueue{ID: row[".id"], Name: row["name"], Target: row["target"], MaxLimit: row["max-limit"], LimitAt: row["limit-at"], Priority: row["priority"], Disabled: boolFromRouterOS(row["disabled"]), Dynamic: boolFromRouterOS(row["dynamic"]), Comment: row["comment"], Raw: row}
}

func simpleQueueSetArgs(set model.SimpleQueueSet) []string {
	args := []string{}
	args = appendArg(args, "name", set.Name)
	args = appendArg(args, "target", set.Target)
	args = appendArg(args, "max-limit", set.MaxLimit)
	args = appendArg(args, "limit-at", set.LimitAt)
	args = appendArg(args, "priority", set.Priority)
	args = appendArg(args, "comment", set.Comment)
	args = appendBoolArg(args, "disabled", set.Disabled)
	return appendExtraArgs(args, set.Extra)
}
