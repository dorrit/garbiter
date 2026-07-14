package garbiter

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dorrit/garbiter/service"
)

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

func appendEnabledArg(args []string, name string, value *bool) []string {
	if value == nil {
		return args
	}
	if *value {
		return append(args, "="+name+"=enabled")
	}
	return append(args, "="+name+"=disabled")
}

func appendIntArg(args []string, name string, value *int) []string {
	if value == nil {
		return args
	}
	return append(args, "="+name+"="+strconv.Itoa(*value))
}

func appendDurationArg(args []string, name string, value *time.Duration) []string {
	if value == nil {
		return args
	}
	return append(args, "="+name+"="+value.String())
}

func appendStringPtrArg(args []string, name string, value *string) []string {
	if value == nil {
		return args
	}
	return append(args, "="+name+"="+*value)
}

func appendExtraArgs(args []string, extra map[string]string) []string {
	existing := make(map[string]struct{}, len(args))
	for _, arg := range args {
		if name := routerOSArgumentName(arg); name != "" {
			existing[name] = struct{}{}
		}
	}

	names := make([]string, 0, len(extra))
	for name := range extra {
		if name == "" || name == ".id" {
			continue
		}
		if _, duplicate := existing[name]; duplicate {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		args = append(args, "="+name+"="+extra[name])
	}
	return args
}

func routerOSArgumentName(arg string) string {
	if len(arg) < 2 || arg[0] != '=' {
		return ""
	}
	index := strings.IndexByte(arg[1:], '=')
	if index < 0 {
		return ""
	}
	return arg[1 : index+1]
}

func boolFromRouterOS(value string) bool {
	return value == "true" || value == "yes"
}

func setIDArgs(id string, args []string) []string {
	return append([]string{"=.id=" + id}, args...)
}

func proplist(fields ...string) string {
	return "=.proplist=" + strings.Join(fields, ",")
}

func validateID(id string) error {
	if id == "" {
		return service.ErrInvalidID
	}
	return nil
}
