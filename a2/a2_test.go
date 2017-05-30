package a2

import (
	"os"
	"testing"
)

func TestReadJSON(t *testing.T) {
	os.Args = []string{"cmd", "input.json"}
    actual := readJSON()

    expected := `{"s":[2, 3], "a < b && a >= c":true}`
    if actual != expected {
        t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
    }
}

func TestReadTokens(t *testing.T) {
	readTokens(readJSON())
}
