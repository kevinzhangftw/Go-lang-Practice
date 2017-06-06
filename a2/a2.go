// package a2
package main

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
	colon
	comma
	boolean
	words
	escapeString
	number
	empty
	leftangle
	rightangle
	amp
	quot
	apos
)

func main() {
	data:=readJSON()
	Token := readTokens(data)
	outHTML := formatHTML(Token, data)
	writeFile(outHTML)
}

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
	commaflag := false
	html := "<!DOCTYPE html><html><head><title>JSON Color</title></head><body>"
	endTag:= "</body></html>"
	body := ""
	for i := 0; i < len(tokenSlice); i++ {
		switch{
		case tokenSlice[i]==delimCurly:
			if string(data[i])=="{" {
				body = body + leftSpan("blue")+ string(data[i]) + "</span><br/>"+ "&nbsp"
			}else{
				body = body + leftSpan("blue")+ "<br/>" +string(data[i]) + "</span>"
			}
			
		case tokenSlice[i]==words:
			body = body + leftSpan("green")+ string(data[i]) + "</span>"
		case tokenSlice[i]==escapeString:
			body = body + leftSpan("PEACHPUFF")+ string(data[i]) + "</span>"
		case tokenSlice[i]==number:
			body = body + leftSpan("cyan")+ string(data[i]) + "</span>"
		case tokenSlice[i]==colon:
			body = body + leftSpan("maroon")+ space()+string(data[i])+ space() + "</span>"
		
		case tokenSlice[i]==delimSquare:
			if string(data[i])=="[" {
				commaflag = true
			}else{
				commaflag = false
			}
			body = body + leftSpan("gray")+ string(data[i]) + "</span>"
		
		case tokenSlice[i]==comma:
			if commaflag == false {
				body = body + leftSpan("red")+ string(data[i]) + "</span><br/>" + "&nbsp"
			}else{
				body = body + leftSpan("red")+ string(data[i])+ space() + "</span>"
			}
			
		case tokenSlice[i]==empty:
			body = body + "&nbsp"
		case tokenSlice[i]==boolean:
			body = body + leftSpan("ORCHID")+ string(data[i]) + "</span>"

		case tokenSlice[i]==leftangle:
			body = body + leftSpan("green")+ "&lt" + "</span>"
		case tokenSlice[i]==rightangle:
			body = body + leftSpan("green")+ "&gt" + "</span>"
		case tokenSlice[i]==amp:
			body = body + leftSpan("green")+ "&amp" + "</span>"
		case tokenSlice[i]==quot:
			body = body + leftSpan("green")+ "&quot" + "</span>"
		case tokenSlice[i]==apos:
			body = body + leftSpan("green")+ "&apos" + "</span>"	
		default:
			// panic("kz says: token not recognized")	
		}
	}


	return html + body + endTag
}

func space() string{
	return "&nbsp"
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
       	case rdat[index]=='<':
       		tokenSlice[index] = leftangle
       	case rdat[index]=='>':
       		tokenSlice[index] = rightangle
       	case rdat[index]=='&':
       		tokenSlice[index] = amp
       	// case rdat[index]=='"':
       	// 	tokenSlice[index] = quot
       	// case rdat[index]=='\'':
       	// 	tokenSlice[index] = apos
    	case rdat[index]== '{', rdat[index]== '}':
       		if strflag==false {
				tokenSlice[index] = delimCurly
			}else{
				tokenSlice[index] = words
			}

       	case rdat[index]== '[', rdat[index]== ']':
       		if strflag==false {
				tokenSlice[index] = delimSquare
			}else{
				tokenSlice[index] = words
			}

       	case rdat[index]==':':
       		if strflag==false {
				tokenSlice[index] = colon
			}else{
				tokenSlice[index] = words
			}

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
