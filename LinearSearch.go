package a1
	
import (
	// "bufio"
    "fmt"
    // "io/ioutil"
    // "os"
	// "strconv"
	// "errors"
	"reflect"
)

//adapted from http://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces-in-go
func linearSearch(x interface{}, lst interface{}) (index int) {
	// xVal := reflect.ValueOf(x)
	// fmt.Println(xVal)

	// lstVal := reflect.ValueOf(lst)
	// if lstVal.Kind() != reflect.Slice {
 //        panic("lst is not a slice")
 //    }
 //    lstInterface := make([]interface{}, lstVal.Len())
 //    for i := 0; i < lstVal.Len(); i++ {
 //    	lstInterface[i] = lstVal.Index(i).Interface    	
 //    }
 //    fmt.Println(lstInterface)
	uselessInt:= 5
	uselessString:="i am useless"
	if reflect.TypeOf(x) == uselessInt.(int) {
		fmt.Println("x is int")	
	}else if reflect.TypeOf(x) == uselessString.(string){
		fmt.Println("x is string")
	}


    //check if x and lst[] is same type



    


	return
}