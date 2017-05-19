package a1
	
import (
	// "bufio"
    // "fmt"
    // "io/ioutil"
    // "os"
	// "strconv"
	// "errors"
	"reflect"
)

//adapted from http://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces-in-go
func linearSearch(x interface{}, lst interface{}) int {
	index := 0	
    newLst := reflect.ValueOf(lst)
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