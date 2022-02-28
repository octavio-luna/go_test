package main

import "testing"

func TestAverageNumber(t *testing.T) {
	tables := []struct {
		s      string
		result float64
	}{
		{"12-as-23-qw", 2},
		{"as-12-as-23-qw", 2},
		{"as-12-as-qw", 1.5},
		{"12-as-qw", 1.5},
		{"12-sdfsd-23-11134", 2},
		{"1fg-12-sdf-113", 1.5},
		{".sad234.3-234-as", 3},
	}
	for _, table := range tables {
		res := averageNumber(table.s)
		if res != table.result {
			t.Errorf("error %s got %f, expected %f", table.s, res, table.result)
		}
	}
}
