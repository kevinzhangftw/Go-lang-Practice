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

func TestValidTime24(t *testing.T) {
	a:= Time24{
		9,30,30,
	}
	if !a.validTime24() {
		t.Errorf("a is validTime24!, fix me")
	}
	b:= Time24{
		25,61,99,
	}
	if b.validTime24() {
		t.Errorf("b is not validTime24!, fix me")
	}
	c:= Time24{
		23,59,99,
	}
	if c.validTime24() {
		t.Errorf("c is not validTime24!, fix me")
	}
	d:= Time24{
		23,59,59,
	}
	if !d.validTime24() {
		t.Errorf("d is validTime24!, fix me")
	}
}

func TestLessThanTime24(t *testing.T) {
	a:= Time24{
		9,45,60,
	}
	b:= Time24{
		10,30,30,
	}
	aComesbeforeb := lessThanTime24(a,b)
	if !aComesbeforeb {
		t.Errorf("a should be before b, but it is not.")
	}

	c:= Time24{
		9,20,30,
	}
	d:= Time24{
		9,30,03,
	}
	cComesbefored := lessThanTime24(c,d)
	if !cComesbefored {
		t.Errorf("c should be before d, but it is not.")
	}
	
	e:= Time24{
		12,59,59,
	}
	f:= Time24{
		12,59,0,
	}
	eComesbeforef := lessThanTime24(e,f)
	if eComesbeforef {
		t.Errorf("e should NOT be before f, but it is.")
	}
}

func TestString(t *testing.T) {
	a:= Time24{
		0,30,30,
	}
	aString := a.String()
	if aString != "00:30:30" {
		t.Errorf("aString should be 00:30:30, not %s", aString)
	}
	b:= Time24{
		19,24,59,
	}
	bString := b.String()
	if bString != "19:24:59" {
		t.Errorf("bString should be 19:24:59, not %s", bString)
	}
	c:= Time24{
		0,0,5,
	}
	cString := c.String()
	if cString != "00:00:05" {
		t.Errorf("cString should be 00:00:05, not %s", cString)
	}
}

func TestMinTime24(t *testing.T) {
	a:= Time24{
		0,30,30,
	}
	b:= Time24{
		19,24,59,
	}
	c:= Time24{
		0,0,5,
	}
	e:= Time24{
		12,59,59,
	}
	f:= Time24{
		12,59,0,
	}
	slice1 := []Time24{a,b,c,e,f}
	emptySlice := []Time24{}
	emptyResult,errEmpty := minTime24(emptySlice)
	slice1Result,errNonEmpty := minTime24(slice1)
	if errEmpty.Error()!="times is empty, please provide a nonempty slice" {
		t.Errorf("emptySlice error object wrong, check error object")
		
	}
	if errNonEmpty.Error()!="nil" {
		t.Errorf("Slice1 error object is wrong, check error object")
	}
	if emptyResult.hour!=0 || emptyResult.minute!=0 || emptyResult.second!=0 {
		t.Errorf("emptySlice allocated non empty Time24 object")
	}
	if slice1Result.hour!=0 || slice1Result.minute!=0 || slice1Result.second!=5 {
		t.Errorf("slice1Result is wrong, should be 0,0,5, but it is %d,%d,%d", slice1Result.hour, slice1Result.minute, slice1Result.second)
	}


}
