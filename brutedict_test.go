package brutedict

import (
	"testing"
)

func TestBruteDict(t *testing.T) {
	bd := New(true, false, false, 3, 3)
	defer bd.Close()
	for {
		id := bd.Id()
		if id != "" {
			t.Error("%d", id)
		}
		if id == "" {
			return
		}
	}
}
