// fline
// program that parses first line of http protocol
//

// ref:
// https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL
//
// V2: read a file with http input

package FLparser

import (
	"testing"
//	"fmt"
//	"log"
)

func TestParser(t *testing.T) {

	// http: 'cmd url proto'
	testLines := make([]SL, 0, 10)
	testLines = append(testLines, []byte("GET / http/1.1\r\n"))
	testLines = append(testLines, []byte("GET /index http/1.1\n"))
	testLines = append(testLines, []byte("GET /index http/1.1\r\n"))

//	fmt.Printf("*** testLines: %d ****\n", len(testLines))
	for i:= 0; i<len(testLines); i++ {
		inp:= testLines[i]
//		fmt.Printf("\nline[%d]: %s", i, inp)
		res, err := ParseFLHttp(inp)
		if err != nil {
			t.Errorf("error -- line: %d ->parseFLHttp: %v\n",i, err)
		} else {
			if len(res.cmd) == 0 {t.Errorf("error -- line: %d ->parseFLHttp no cmd\n",i)}
		}
	}
//	fmt.Println("*** success ***")
}

