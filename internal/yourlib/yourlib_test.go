package yourlib_test

import "testing"

func skipShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
}

func TestFast(t *testing.T) {

}

func TestSlow(t *testing.T) {
	skipShort(t)
}
