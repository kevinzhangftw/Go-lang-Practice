package a1
	
import (
    "fmt"
	"reflect"
)

//adapted from http://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces-in-go
func linearSearch(x interface{}, lst interface{}) int {
	index := 0	
    newLst := reflect.ValueOf(lst)
    //check type
    if reflect.TypeOf(x) != reflect.TypeOf(newLst.Index(0).Interface()) {
    	fmt.Println("panic() commented out: x and lst not the same type")
    	// panic("x and lst not the same type")
    }

    for i := 0; i < newLst.Len(); i++ {
    	if x == newLst.Index(i).Interface() {
    		index = i
    		return index
    	} 	
    }
//otherwise here x is not found in lst
    index = -1
	return index
}