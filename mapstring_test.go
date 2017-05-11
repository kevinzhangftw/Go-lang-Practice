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
	equalTest1 := reflect.DeepEqual(firstWordCount, rightFirstWordCount)
	if !equalTest1 {
		t.Errorf("Unexpected word count , it is not %s.", firstWordCount)
	}
	//adapted from stackoverflow http://stackoverflow.com/questions/18208394/testing-equivalence-of-maps-golang
	
	secondWordCount := countStrings("test2.txt")
	rightsecondWordCount := map[string]int{
		"The":1, 
		"the":4,
		"pair":1, 
	}
	equalTest2 := reflect.DeepEqual(secondWordCount, rightsecondWordCount)
	if !equalTest2 {
		t.Errorf("Unexpected word count , it is not %s.", secondWordCount)
	}
	
	thirdWordCount := countStrings("test3.txt")
	rightthirdWordCount := make(map[string]int)
	equalTest3 := reflect.DeepEqual(thirdWordCount, rightthirdWordCount)
	if !equalTest3 {
		t.Errorf("Unexpected word count , it is not %s.", thirdWordCount)
	}

}
