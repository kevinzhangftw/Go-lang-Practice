package a1
	
import (
	// "bufio"
    // "fmt"
    // "io/ioutil"
    // "os"
)

type Time24 struct {
    hour, minute, second uint8
}
// 0 <= hour < 24
// 0 <= minute < 60
// 0 <= second < 60

func equalsTime24(a Time24, b Time24) bool {
    equal := false
    if a.hour==b.hour && a.minute==b.minute && a.second==b.second {
    	equal = true
    }
    return equal
}

func lessThanTime24(a Time24, b Time24) bool {
	aComesbeforeb := false


	return aComesbeforeb
}