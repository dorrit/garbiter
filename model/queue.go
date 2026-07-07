package model

type SimpleQueue struct {
	ID       string
	Name     string
	Target   string
	MaxLimit string
	LimitAt  string
	Priority string
	Disabled bool
	Dynamic  bool
	Comment  string
	Raw      map[string]string
}

type SimpleQueueSet struct {
	Name     string
	Target   string
	MaxLimit string
	LimitAt  string
	Priority string
	Comment  string
	Disabled *bool
	Extra    map[string]string
}
