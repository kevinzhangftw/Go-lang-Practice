package a1
	
import (
	"bufio"
    "fmt"
    "io/ioutil"
    "os"
)

//adapted from https://gobyexample.com/reading-files
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func countStrings(filename string) map[string]int {
	wordcount := map[string]int{
		"nil":0,
	}

	dat, err := ioutil.ReadFile(filename)
    check(err)
    fmt.Print(string(dat))

// adapted from https://www.dotnetperls.com/file-go
	f, err := os.Open(filename)
    check(err)
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
    	//each line is a word
        line := scanner.Text()
        //if it is in map lookup ok, update value
        //else it is not, update the string
        fmt.Println(line)
    }

	return wordcount
}