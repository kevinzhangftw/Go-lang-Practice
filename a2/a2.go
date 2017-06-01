package a2

import (
	// "encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	// "io"
	// "log"
	// "strings"
)

type Token interface{}

type TokenType uint8

const (
	delimCurly TokenType = iota
	delimSquare 
	quote
	comma
	boolean
	words
	escapeString
	number
)

// adapted from https://stackoverflow.com/questions/35080109/golang-how-to-read-input-filename-in-go
func readJSON() string {
	if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return "Missing parameter, provide file name!"
    }

    data, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Can't read file:", os.Args[1])
        panic(err)
    }
   
    return string(data)
}

func readTokens(data string) []Token{
	tokenSlice := make([]Token, len(data))
 
	rdat := []rune(data)
	for index := range rdat {
		switch {
    	case rdat[index]== '{', rdat[index]== '}':
       		tokenSlice[index] = delimCurly

       	case rdat[index]== '[', rdat[index]== ']':
       		tokenSlice[index] = delimSquare

       	case rdat[index]==':':
       		tokenSlice[index] = quote

       	case rdat[index]==',':
       		tokenSlice[index] = comma

       	case rdat[index]=='0',rdat[index]=='1',rdat[index]=='2',rdat[index]=='3':
       		tokenSlice[index] = number
       	case rdat[index]=='4',rdat[index]=='5',rdat[index]=='6',rdat[index]=='7':
       		tokenSlice[index] = number
       	case rdat[index]=='8',rdat[index]=='9',rdat[index]=='+',rdat[index]=='-':
       		tokenSlice[index] = number
       	case rdat[index]=='e',rdat[index]=='E',rdat[index]=='.':
       		tokenSlice[index] = number

       	case rdat[index]=='t',rdat[index]=='r',rdat[index]=='u',rdat[index]=='e':
       		tokenSlice[index] = boolean

    	default:
       		panic("Kz says: invalid JSON")
    	}
    }

	return tokenSlice
}
