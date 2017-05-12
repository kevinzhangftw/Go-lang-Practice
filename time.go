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
	if a.hour<b.hour {
		aComesbeforeb = true
		return aComesbeforeb
	}else if a.hour>b.hour {
		return aComesbeforeb
	}else{
		//hour is same, check minute
		if a.minute<b.minute {
			aComesbeforeb = true
			return aComesbeforeb
		}else if a.minute>b.minute {
			return aComesbeforeb
		}else{
			// minute is same check second
			if a.second<b.second {
				aComesbeforeb = true
				return aComesbeforeb
			}else if a.second>b.second {
				return aComesbeforeb
			}else{
				return aComesbeforeb
			}
		}
	}
	return aComesbeforeb
}