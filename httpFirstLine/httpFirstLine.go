// fline
// program that parses first line of http protocol
//

// ref:
// https://developer.mozilla.org/en-US/docs/Learn/Common_questions/Web_mechanics/What_is_a_URL
//
// V2: read a file with http input

package FLparser

import (
	"fmt"
//	"log"
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


func ParseFLHttp(inp []byte)(res resHttp, err error) {

	state:=0
	ist := 0
	for i:=0; i<len(inp); i++ {

//		fmt.Printf("state: %d ist
		switch state {
		// check ' '
		case 0:
			if inp[i] == ' ' {
				res.cmd = inp[:i]
				state = 1
				ist = i+1
			}
			if inp[i] == '\n' {
				return res, fmt.Errorf("cmd not found!")
			}
		// parse url
		case 1:
			if inp[i] == ' ' {
				res.url = inp[ist:i]
				state = 2
				ist = i+1
			}
			if inp[i] == '\n' {
				return res, fmt.Errorf("url not found!")
			}
		// parse http proto
		case 2:
			if inp[i] == '\r' {
				res.proto = inp[ist:i]
				switch inp[i-1] {
				case '0':
					res.prot1 = false
				case '1':
					res.prot1 = true
				default:
					return res, fmt.Errorf("invalid http proto")
				}
				state = 3
			}
			if inp[i] == '\n' {
				return res, fmt.Errorf("eol has no cr!")
			}

		case 3:
			if inp[i] != '\n' {
				return res, fmt.Errorf("eol has no lf!")
			}

		default:
			return res, fmt.Errorf("invalid state %d", state)
		}
	}

	return res, nil
}

func PrintRes(res resHttp) {
	fmt.Println("****** parse Http ******")
	fmt.Printf("cmd:     %s\n", res.cmd)
	fmt.Printf("url:     %s\n", res.url)
	fmt.Printf("proto:   %s\n", res.proto)
	fmt.Printf("http1.1: %t\n", res.prot1)
	fmt.Println("**** end parse Http ****")
}
