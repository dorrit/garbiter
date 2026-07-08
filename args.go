package garbiter

import "github.com/dorrit/garbiter/service"

func checkService(svc interface{}) error {
	if svc == nil {
		return service.ErrNotConnected
	}
	return nil
}

func appendArg(args []string, name, value string) []string {
	if value == "" {
		return args
	}
	return append(args, "="+name+"="+value)
}

func appendBoolArg(args []string, name string, value *bool) []string {
	if value == nil {
		return args
	}
	if *value {
		return append(args, "="+name+"=yes")
	}
	return append(args, "="+name+"=no")
}

func appendExtraArgs(args []string, extra map[string]string) []string {
	for name, value := range extra {
		args = appendArg(args, name, value)
	}
	return args
}

func boolFromRouterOS(value string) bool {
	return value == "true" || value == "yes"
}

func setIDArgs(id string, args []string) []string {
	return append([]string{"=.id=" + id}, args...)
}

func validateID(id string) error {
	if id == "" {
		return service.ErrInvalidID
	}
	return nil
}
