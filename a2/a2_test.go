package a2

import (
	"os"
	"testing"
	"reflect"
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
	outEmpty := readTokens(``)
	if len(outEmpty) != 0 {
		t.Errorf("Test failed, expected: 0, got:  '%d'", len(outEmpty))	
	}
	outShorty := readTokens(`{}`)
	expectShorty := make([]Token, 2)
	expectShorty[0] = delimCurly
	expectShorty[1] = delimCurly
	if !reflect.DeepEqual(outShorty, expectShorty) {
		t.Errorf("Test failed, expected: '%v', got:  '%s'", expectShorty, outShorty)
	}
	
	outSimple := readTokens(`{:[,]}`)
	expectSimple := make([]Token, 6)
	expectSimple[0] = delimCurly
	expectSimple[1] = quote
	expectSimple[2] = delimSquare
	expectSimple[3] = comma
	expectSimple[4] = delimSquare
	expectSimple[5] = delimCurly
	if !reflect.DeepEqual(outSimple, expectSimple) {
		t.Errorf("Test failed, expected: '%v', got:  '%s'", expectSimple, outSimple)
	}
}
