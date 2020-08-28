package server

import "testing"

func Test_World_CreateWorld(t *testing.T) {
	for i := 0; i < 16; i++ {
		CreateWorld([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i"})
	}
}
