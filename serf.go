package main

type serfFilter struct {
	Tags   map[string]string
	Status string
	Name   string
}

func (sf1 *serfFilter) Compare(sf2 serfFilter) bool {
	return true
}
