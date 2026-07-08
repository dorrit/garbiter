package model

type LogEntry struct {
	ID      string
	Time    string
	Topics  string
	Message string
	Raw     map[string]string
}
