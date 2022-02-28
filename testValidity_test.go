package main

import (
	"testing"
)

func TestTestValidity(t *testing.T) {
	tables := []struct {
		s      string
		result bool
	}{
		{"12-as-23-qw", true},
		{"as-12-as-23-qw", false},
		{"as-12-as-qw", false},
		{"12-as-qw", false},
		{"12-sdfsd-23-1234", false},
		{"1fg-12-sdf-23", false},
		{".sad234.4-234-as", false},
	}

	for _, table := range tables {
		res := testValidity(table.s)
		if res != table.result {
			t.Errorf("error %s got %t, expected %t", table.s, res, table.result)
		}
	}
}
