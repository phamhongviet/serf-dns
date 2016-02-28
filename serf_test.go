package main

import (
	"testing"
)

func TestSerfFilterCompare(t *testing.T) {
	sf1 := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"a": "65",
			"b": "66",
			"c": "67",
		},
	}
	sf2 := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"a": "65",
			"b": "66",
			"c": "67",
		},
	}
	sf3 := serfFilter{
		Name:   "",
		Status: "alive",
		Tags: map[string]string{
			"a": "65",
			"b": "66",
		},
	}
	if sf1.Compare(sf2) != true {
		t.Errorf("SerfFilter Compare return unexpected result.")
	}
	if sf1.Compare(sf3) != false {
		t.Errorf("SerfFilter Compare return unexpected result.")
	}
}
