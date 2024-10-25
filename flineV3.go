// fline
// program that parses first line of http protocol
//

// ref:
// https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL
//
// V2: read a file with http input
// V3: test library

package main

import (
	"fmt"
//	"log"

	pars "server/http/httpParser/httpFirstLine"
//server/http/httpParser/httpFirstLine
)

type resHttp struct {
	cmd []byte
	url []byte
	kv map[string] string
	https bool
	domain []byte
	anchor []byte
	proto []byte
	prot1	bool
}

type SL []byte

func main() {

	// http: 'cmd url proto'
	testLines := make([]SL, 0, 10)
	testLines = append(testLines, []byte("GET / http/1.1\r\n"))
	testLines = append(testLines, []byte("GET /index http/1.1\n"))
	testLines = append(testLines, []byte("GET /index http/1.1\r\n"))

	fmt.Printf("*** testLines: %d ****\n", len(testLines))
	for i:= 0; i<len(testLines); i++ {
		inp:= testLines[i]
		fmt.Printf("\nline[%d]: %s", i, inp)
		res, err := pars.ParseFLHttp(inp)
		if err != nil {
			fmt.Printf("error -- line: %d ->parseFLHttp: %v\n",i, err)
		} else {
			pars.PrintRes(res)
		}
	}
	fmt.Println("*** success ***")
}

