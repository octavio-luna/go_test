package main

import "testing"

func TestStoryStats(t *testing.T) {
	tables := []struct {
		s       string
		short   string
		long    string
		average float64
		words   []string
	}{
		{"12-as-23-qw", "as", "as", 2.0, []string{"as", "qw"}},
		{"as-12-s-23-qwwaaa", "s", "qwwaaa", 3, []string{"as", "s", "qwwaaa"}},
		{"as", "as", "as", 2, []string{"as"}},
		{"", "", "", 0, []string{}},
		{"12", "", "", 0, []string{}},
	}

	for _, table := range tables {
		s, l, a, w := storyStats(table.s)
		if s != table.short {
			t.Errorf("Error %s, expected %s, got %s", table.s, table.short, s)
		}
		if l != table.long {
			t.Errorf("Error %s, expected %s, got %s", table.s, table.long, l)
		}
		if a != table.average {
			t.Errorf("Error %s, expected %f, got %f", table.s, table.average, a)
		}
		for i := 0; i < len(w); i++ {
			if w[i] != table.words[i] {
				t.Errorf("Error %s, row %v expected %s, got %s", table.s, i, table.words[i], w[i])
			}
		}
	}
}
