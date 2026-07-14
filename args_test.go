package garbiter

import (
	"reflect"
	"testing"
)

func TestAppendExtraArgsSupportsClearingAndStableOrdering(t *testing.T) {
	args := appendExtraArgs(
		[]string{"=.id=*1", "=comment=typed"},
		map[string]string{
			"z-field": "last",
			"comment": "duplicate",
			"a-field": "",
			".id":     "*2",
		},
	)

	want := []string{"=.id=*1", "=comment=typed", "=a-field=", "=z-field=last"}
	if !reflect.DeepEqual(args, want) {
		t.Fatalf("args = %#v, want %#v", args, want)
	}
}
