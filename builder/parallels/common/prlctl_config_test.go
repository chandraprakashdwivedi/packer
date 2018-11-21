package common

import (
	"reflect"
	"testing"
)

func TestPrlctlConfigPrepare_Prlctl(t *testing.T) {
	// Test with empty
	c := new(PrlctlConfig)
	errs := c.Prepare(testConfigTemplate(t))
	if len(errs) > 0 {
		t.Fatalf("err: %#v", errs)
	}

	// Test with a good one
	c = new(PrlctlConfig)
	c.Prlctl = []string{
		"foo", "bar", "baz",
	}
	errs = c.Prepare(testConfigTemplate(t))
	if len(errs) > 0 {
		t.Fatalf("err: %#v", errs)
	}

	expected := []string{
		"foo", "bar", "baz",
	}

	if !reflect.DeepEqual(c.Prlctl, expected) {
		t.Fatalf("bad: %#v", c.Prlctl)
	}
}
