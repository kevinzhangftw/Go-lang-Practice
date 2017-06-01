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
	
	outSimple := readTokens(`{:[-52.01e-355]}`)
	expectSimple := make([]Token, 16)
	expectSimple[0] = delimCurly
	expectSimple[1] = quote
	expectSimple[2] = delimSquare
	expectSimple[3] = number
	expectSimple[4] = number
	expectSimple[5] = number
	expectSimple[6] = number
	expectSimple[7] = number
	expectSimple[8] = number
	expectSimple[9] = number
	expectSimple[10] = number
	expectSimple[11] = number
	expectSimple[12] = number
	expectSimple[13] = number
	expectSimple[14] = delimSquare
	expectSimple[15] = delimCurly
	if !reflect.DeepEqual(outSimple, expectSimple) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectSimple, outSimple)
	}

	outBool := readTokens(`{:true}`)
	expectBool := make([]Token, 7)
	expectBool[0] = delimCurly
	expectBool[1] = quote
	expectBool[2] = boolean
	expectBool[3] = boolean
	expectBool[4] = boolean
	expectBool[5] = boolean
	expectBool[6] = delimCurly
	if !reflect.DeepEqual(outBool, expectBool) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectBool, outBool)
	}
}
