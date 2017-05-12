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

}
