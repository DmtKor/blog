package commander

import "testing"

func TestCommander(t *testing.T) {
	c := Commander{}
	c.InitServer(initstr_temp, ":50050")
}