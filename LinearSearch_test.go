package a1
	
import (
	"testing"
	"fmt"
)

func testLinearSearch(t *testing.T) {
	test1 := linearSearch(5, []int{4, 2, -1, 5, 0})
	fmt.Println(test1)
	if test1 != 3 {
		t.Errorf("test1 should return 3, instead it returns %v", test1)
	}

	
}