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
	empty
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
 	strflag := false
 	// escflag := false
 	
	rdat := []rune(data)
	for index := 0; index < len(rdat); index++ {
		switch {
		case rdat[index]==' ':
       		tokenSlice[index] = empty

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
       	case rdat[index]=='E',rdat[index]=='.':
       		tokenSlice[index] = number

		case rdat[index]=='t',rdat[index]=='r',rdat[index]=='u',rdat[index]=='n':
			if strflag==false {
				tokenSlice[index] = boolean
			}else{
				tokenSlice[index] = words
			}
       		
		case rdat[index]=='f',rdat[index]=='a',rdat[index]=='l',rdat[index]=='s':
       		if strflag==false {
				tokenSlice[index] = boolean
			}else{
				tokenSlice[index] = words
			}
		
       	case rdat[index]=='e':
       		if strflag ==false {
       			if tokenSlice[index-1]==number {
       				tokenSlice[index]=number
       			}else{
       				tokenSlice[index]=boolean
       			}	
       		}else{
       			tokenSlice[index]= words
       		}

       	case rdat[index]=='"':
       		tokenSlice[index]=words
       		if strflag==false {
       			strflag= true
       		}else{
       			strflag= false
       		}

       	case rdat[index]=='\\':
       		// escflag = !escflag
       		if rdat[index+1]!='u' {
       			tokenSlice[index]= escapeString
       			tokenSlice[index+1]=escapeString
       			index++	
       		}else{
       			tokenSlice[index]= escapeString
       			tokenSlice[index+1]=escapeString
       			tokenSlice[index+2]= escapeString
       			tokenSlice[index+3]=escapeString
       			tokenSlice[index+4]= escapeString
       			tokenSlice[index+5]=escapeString
       			index+=5
       		}
       		
    	default:
    		if strflag == true {
    			tokenSlice[index]= words
    		}
    	}
    }


	return tokenSlice
}

func formatHTMl(tokenSlice []Token) string{
	html := "fake html"

	return html
}
