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
 
	runedata := []rune(data)
	for index := range runedata {

		switch {
    	case runedata[index]== '{', runedata[index]== '}':
       		tokenSlice[index] = delimCurly

       	case runedata[index]== '[', runedata[index]== ']':
       		tokenSlice[index] = delimSquare

       	case runedata[index]==':':
       		tokenSlice[index] = quote

       	case runedata[index]==',':
       		tokenSlice[index] = comma

    	default:
       		panic("Kz says: invalid JSON")
    	}
    }

	return tokenSlice
}
