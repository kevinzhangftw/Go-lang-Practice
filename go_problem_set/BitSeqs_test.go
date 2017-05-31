package a1
	
import (
    // "fmt"
	"reflect"
	"testing"
)

func TestAllBitSeqs(t *testing.T) {
	test0 := allBitSeqs(0)
	emptySlice := make([][]int, 0)
	if reflect.DeepEqual(test0, emptySlice) == false {
		t.Errorf("test0 should return empty slice, instead it returns %v", test0)
	}

	slice1 := [][]int{}
	slice1 = append(slice1, []int{0})
	slice1 = append(slice1, []int{1})
	test1 := allBitSeqs(1)
	if reflect.DeepEqual(test1, slice1) == false {
		t.Errorf("test1 should return [[0] [1]], instead it returns %v", test1)
	}	

	slice2 := [][]int{}
	slice2 = append(slice2, []int{0, 0})
	slice2 = append(slice2, []int{0, 1})
	slice2 = append(slice2, []int{1, 0})
	slice2 = append(slice2, []int{1, 1})
	test2 := allBitSeqs(2)
	if reflect.DeepEqual(test2, slice2) == false {
		t.Errorf("test2 should return [[0 0] [0 1] [1 0] [1 1]], instead it returns %v", test2)
	}
}