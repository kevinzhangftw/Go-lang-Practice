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

func formatHTML(tokenSlice []Token, data string) string{
	html := "<!DOCTYPE html><html><head><title>JSON Color</title></head><body>"
	endTag:= "</body></html>"
	body := ""
	for i := 0; i < len(tokenSlice); i++ {
		switch{
		case tokenSlice[i]==delimCurly:
			if string(data[i])=="{" {
				body = body + leftSpan("blue")+ string(data[i]) + "</span><br/>"
			}else{
				body = body + leftSpan("blue")+ "<br/>" +string(data[i]) + "</span>"
			}
			
		case tokenSlice[i]==words:
			body = body + leftSpan("green")+ string(data[i]) + "</span>"
		case tokenSlice[i]==escapeString:
			body = body + leftSpan("PEACHPUFF")+ string(data[i]) + "</span>"
		case tokenSlice[i]==number:
			body = body + leftSpan("cyan")+ string(data[i]) + "</span>"
		case tokenSlice[i]==quote:
			body = body + leftSpan("maroon")+ string(data[i]) + "</span>"
		case tokenSlice[i]==delimSquare:
			body = body + leftSpan("gray")+ string(data[i]) + "</span>"
		case tokenSlice[i]==comma:
			body = body + leftSpan("red")+ string(data[i]) + "</span><br/>"
		case tokenSlice[i]==empty:
			body = body + "&nbsp"
		case tokenSlice[i]==boolean:
			body = body + leftSpan("ORCHID")+ string(data[i]) + "</span>"

		default:
			panic("kz says: token not recognized")	
		}
	}


	return html + body + endTag
}

// adapted from https://stackoverflow.com/questions/35333302/how-to-write-the-output-of-this-statement-into-a-file-in-golang
func writeFile(html string) {
	file, fileErr := os.Create("json.html")
	if fileErr != nil {
    	fmt.Println(fileErr)
    	return
	}
	fmt.Fprintf(file, "%v\n", html)
}

func leftSpan(color string) string{
	return `<span style="color:` + color + `">`
}


func readTokens(data string) []Token{
	tokenSlice := make([]Token, len(data))
 	strflag := false
 	
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
       		if strflag==false {
				tokenSlice[index] = comma
			}else{
				tokenSlice[index] = words
			}

       	case rdat[index]=='0',rdat[index]=='1',rdat[index]=='2',rdat[index]=='3':
       		if strflag==false {
				tokenSlice[index] = number
			}else{
				tokenSlice[index] = words
			}
       	case rdat[index]=='4',rdat[index]=='5',rdat[index]=='6',rdat[index]=='7':
       		if strflag==false {
				tokenSlice[index] = number
			}else{
				tokenSlice[index] = words
			}
       	case rdat[index]=='8',rdat[index]=='9',rdat[index]=='+',rdat[index]=='-':
       		if strflag==false {
				tokenSlice[index] = number
			}else{
				tokenSlice[index] = words
			}
       	case rdat[index]=='E',rdat[index]=='.':
       		if strflag==false {
				tokenSlice[index] = number
			}else{
				tokenSlice[index] = words
			}

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
