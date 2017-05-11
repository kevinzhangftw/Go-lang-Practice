package a1

import "testing"
import "reflect"

func TestCountStrings(t *testing.T) {

	firstWordCount := countStrings("test1.txt")

	rightFirstWordCount := map[string]int{
		"The":1, 
		"the":1, 
		"big":3, 
		"dog":1, 
		"ate":1, 
		"apple":1,
	}

	equal := reflect.DeepEqual(firstWordCount, rightFirstWordCount)
	if !equal {
		t.Errorf("Unexpected word count , it is not %s.", firstWordCount)
	}
	//adapted from stackoverflow http://stackoverflow.com/questions/18208394/testing-equivalence-of-maps-golang

}
