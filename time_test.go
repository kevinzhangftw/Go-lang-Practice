package a1
	
import (
	"testing"
)

func TestEqualsTime24(t *testing.T) {
	a:= Time24{
		9,30,30,
	}
	b:= Time24{
		9,30,30,
	}
	abTimeIsEqual := equalsTime24(a,b)
	if !abTimeIsEqual {
		t.Errorf("time for a and b should be equal, but it is not.")
	}

	c:= Time24{
		9,30,30,
	}
	d:= Time24{
		9,30,3,
	}
	cdTimeIsEqual := equalsTime24(c,d)
	if cdTimeIsEqual {
		t.Errorf("time for c and d should NOT be equal, but it is.")
	}
}
