package main

import "testing"

func TestWholeStory(t *testing.T) {
	tables := []struct {
		s      string
		result string
	}{
		{"12-as-23-qw", " as qw"},
		{"as-12-as-23-qw", "as as qw"},
		{"", ""},
		{" as-qw", " as qw"},
		{"12-sdfsd-23-11134", " sdfsd"},
		{"1fg-12-sdf-113", " fg sdf"},
		{".sad234.3-234-as", " sad as"},
	}
	for _, table := range tables {
		res := wholeStory(table.s)
		if res != table.result {
			t.Errorf("error %s got %s, expected %s", table.s, res, table.result)
		}
	}
}
