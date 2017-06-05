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

func TestFormatHTMl(t *testing.T) {
	
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

	outBool := readTokens(`{null}`)
	expectBool := make([]Token, 6)
	expectBool[0] = delimCurly
	expectBool[1] = boolean
	expectBool[2] = boolean
	expectBool[3] = boolean
	expectBool[4] = boolean
	expectBool[5] = delimCurly
	if !reflect.DeepEqual(outBool, expectBool) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectBool, outBool)
	}

	outStr := readTokens(`{"ue"}`)
	expectStr := make([]Token, 6)
	expectStr[0] = delimCurly
	expectStr[1] = words
	expectStr[2] = words
	expectStr[3] = words
	expectStr[4] = words
	expectStr[5] = delimCurly
	if !reflect.DeepEqual(outStr, expectStr) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectStr, outStr)
	}

	outEsc := readTokens(`{"\n"}`)
	expectEsc := make([]Token, 6)
	expectEsc[0] = delimCurly
	expectEsc[1] = words
	expectEsc[2] = escapeString
	expectEsc[3] = escapeString
	expectEsc[4] = words
	expectEsc[5] = delimCurly
	if !reflect.DeepEqual(outEsc, expectEsc) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectEsc, outEsc)
	}

	outEscu := readTokens(`{"\u1234"}`)
	expectEscu := make([]Token, 10)
	expectEscu[0] = delimCurly
	expectEscu[1] = words
	expectEscu[2] = escapeString
	expectEscu[3] = escapeString
	expectEscu[4] = escapeString
	expectEscu[5] = escapeString
	expectEscu[6] = escapeString
	expectEscu[7] = escapeString
	expectEscu[8] = words
	expectEscu[9] = delimCurly
	if !reflect.DeepEqual(outEscu, expectEscu) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectEscu, outEscu)
	}

	outLong := readTokens(`{"s":[2,3],"a<b&&a>=c":true}`)
	expectLong := make([]Token, 28)
	expectLong[0] = delimCurly
	expectLong[1] = words
	expectLong[2] = words
	expectLong[3] = words
	expectLong[4] = quote
	expectLong[5] = delimSquare
	expectLong[6] = number
	expectLong[7] = comma
	expectLong[8] = number
	expectLong[9] = delimSquare
	expectLong[10] = comma
	expectLong[11] = words
	expectLong[12] = words
	expectLong[13] = words
	expectLong[14] = words
	expectLong[15] = words
	expectLong[16] = words
	expectLong[17] = words
	expectLong[18] = words
	expectLong[19] = words
	expectLong[20] = words
	expectLong[21] = words
	expectLong[22] = quote
	expectLong[23] = boolean
	expectLong[24] = boolean
	expectLong[25] = boolean
	expectLong[26] = boolean
	expectLong[27] = delimCurly
	if !reflect.DeepEqual(outLong, expectLong) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectLong, outLong)
	}

	outLongs := readTokens(`{"s":[2, 3], "a < b && a >= c":true}`)
	expectLongs := make([]Token, 36)
	expectLongs[0] = delimCurly
	expectLongs[1] = words
	expectLongs[2] = words
	expectLongs[3] = words
	expectLongs[4] = quote
	expectLongs[5] = delimSquare
	expectLongs[6] = number
	expectLongs[7] = comma
	expectLongs[8] = empty
	expectLongs[9] = number
	expectLongs[10] = delimSquare
	expectLongs[11] = comma
	expectLongs[12] = empty
	expectLongs[13] = words
	expectLongs[14] = words
	expectLongs[15] = empty
	expectLongs[16] = words
	expectLongs[17] = empty
	expectLongs[18] = words
	expectLongs[19] = empty
	expectLongs[20] = words
	expectLongs[21] = words
	expectLongs[22] = empty
	expectLongs[23] = words
	expectLongs[24] = empty
	expectLongs[25] = words
	expectLongs[26] = words
	expectLongs[27] = empty
	expectLongs[28] = words
	expectLongs[29] = words
	expectLongs[30] = quote
	expectLongs[31] = boolean
	expectLongs[32] = boolean
	expectLongs[33] = boolean
	expectLongs[34] = boolean
	expectLongs[35] = delimCurly
	if !reflect.DeepEqual(outLongs, expectLongs) {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expectLongs, outLongs)
	}

}
