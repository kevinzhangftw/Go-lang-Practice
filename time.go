package a1
	
import (
	// "bufio"
    // "fmt"
    // "io/ioutil"
    // "os"
	"strconv"
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

//adapted from http://stackoverflow.com/questions/19223277/golang-convert-uint8-to-string
func (t Time24) String() string {
	timeString:=""
	if t.hour>9 && t.minute>9 && t.second>9 {
		timeString=strconv.Itoa(int(t.hour))+":"+strconv.Itoa(int(t.minute))+":"+strconv.Itoa(int(t.second))
	}else if t.hour<=9 && t.minute>9 && t.second>9 {
		timeString="0"+strconv.Itoa(int(t.hour))+":"+strconv.Itoa(int(t.minute))+":"+strconv.Itoa(int(t.second))
	}else if t.hour<=9 && t.minute<=9 && t.second>9 {
		timeString="0"+strconv.Itoa(int(t.hour))+":0"+strconv.Itoa(int(t.minute))+":"+strconv.Itoa(int(t.second))
	}else if t.hour<=9 && t.minute<=9 && t.second<=9 {
		timeString="0"+strconv.Itoa(int(t.hour))+":0"+strconv.Itoa(int(t.minute))+":0"+strconv.Itoa(int(t.second))
	}else if t.hour>9 && t.minute<=9 && t.second>9 {
		timeString=strconv.Itoa(int(t.hour))+":0"+strconv.Itoa(int(t.minute))+":"+strconv.Itoa(int(t.second))
	}else if t.hour>9 && t.minute<=9 && t.second<=9 {
		timeString=strconv.Itoa(int(t.hour))+":0"+strconv.Itoa(int(t.minute))+":0"+strconv.Itoa(int(t.second))
	}else if t.hour>9 && t.minute>9 && t.second<=9 {
		timeString=strconv.Itoa(int(t.hour))+":"+strconv.Itoa(int(t.minute))+":0"+strconv.Itoa(int(t.second))
	}
    return timeString
}

