package showhex

import "testing"

func TestShowHex(t *testing.T) {
	ShowHex([]byte("Hello World, It's a wonderful day."))
}
