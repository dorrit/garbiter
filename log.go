package garbiter

import (
	"github.com/dorrit/garbiter/model"
	"github.com/dorrit/garbiter/service"
)

type LogAPI struct {
	svc model.Transport
}

func (api *LogAPI) Print() ([]model.LogEntry, error) {
	if api == nil || api.svc == nil {
		return nil, service.ErrNotConnected
	}
	rows, err := api.svc.RunList("/log/print", proplist(".id", "time", "topics", "message"))
	if err != nil {
		return nil, err
	}
	items := make([]model.LogEntry, 0, len(rows))
	for _, row := range rows {
		items = append(items, model.LogEntry{ID: row[".id"], Time: row["time"], Topics: row["topics"], Message: row["message"], Raw: row})
	}
	return items, nil
}
