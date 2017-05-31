package a1
	
import (
	"bufio"
    // "fmt"
    // "io/ioutil"
    "os"
)

//adapted from https://gobyexample.com/reading-files
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func countStrings(filename string) map[string]int {
	wordcount := make(map[string]int)
// adapted from https://www.dotnetperls.com/file-go
	f, err := os.Open(filename)
    check(err)
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        eachWord := scanner.Text()
        // adapted from http://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
        if _, ok := wordcount[eachWord]; ok {
        	wordcount[eachWord]++ 	
        }else{
        	wordcount[eachWord] = 1
        }        
    }
	return wordcount
}