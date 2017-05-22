package a1
	
import (
	"testing"
	// "fmt"
)

func TestLinearSearch(t *testing.T) {
	test1 := linearSearch(5, []int{4, 2, -1, 5, 0})	
	if test1 != 3 {
		t.Errorf("test1 should return 3, instead it returns %v", test1)
	}
	test2 := linearSearch("egg", []string{"cat", "nose", "egg"})	
	if test2 != 2 {
		t.Errorf("test2 should return 2, instead it returns %v", test2)
	}
	test3 := linearSearch(3, []int{4, 2, -1, 5, 0})	
	if test3 != -1 {
		t.Errorf("test3 should not be found, instead it returns %v", test3)
	}
	test4 := linearSearch("up", []string{"cat", "nose", "egg"})	
	if test4 != -1 {
		t.Errorf("test4 should not be found, instead it returns %v", test4)
	}
	var b int =5
	var a *int = &b
	test5 := linearSearch(a, []*int{a})	
	if test5 != 0 {
		t.Errorf("test5 should be 2, instead it returns %v", test5)
	}
	test6 := linearSearch(5, []string{"cat", "nose", "egg"})	
	if test6 != -1 {
		t.Errorf("test6 should panic %v", test6)
	}	
}