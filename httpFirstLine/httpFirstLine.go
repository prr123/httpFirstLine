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
	path []byte
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
	key := ""
	val := ""
	kv := make(map[string]string)

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
			switch inp[i] {
			case '/':
				if inp[i+1] == '/' {
					i = i+1
				} else {
					res.domain = inp[ist:i]
					state = 2
					ist = i+1
				}
			case ' ':
				res.path = inp[ist:i]
				state = 3
				ist = i+1

			case '\r', '\n':
				return res, fmt.Errorf("cannot parse domain -- no '/' char!")
			default:
			}

		case 2:
			switch inp[i] {
			case ' ':
				res.path = inp[ist:i]
				state = 3
				ist = i+1

			case '?':
				res.path = inp[ist:i]
				state = 12
				ist = i+1

			case '#':
				res.path = inp[ist:i]
				state = 15
				ist = i+1

			case '\n':
				return res, fmt.Errorf("cannot parse path -- no ws char!")

			default:
			}

		// parse http proto
		case 3:
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
				state = 4
			}
			if inp[i] == '\n' {
				return res, fmt.Errorf("eol has no cr!")
			}

		// check ending
		case 4:
			if inp[i] != '\n' {
				return res, fmt.Errorf("eol has no lf!")
			}

		//parse kv
		case 12:
			switch inp[i] {
				case '=':
					key = string(inp[ist:i])
					state = 13
					ist = i+1

				case '\n','#',' ', '&':
					return res, fmt.Errorf("state 11 no key!")
				default:
			}
		// parse value
		case 13:
			switch inp[i] {
				// another kv will follow
				case '&':
					val = string(inp[ist:i])
					kv[key] = val
					state = 12
					ist = i+1

				// end of kv; parse protocol next
				case ' ':
					val = string(inp[ist:i])
					kv[key] = val
					state = 3
					ist = i+1

				// end of kv followed by anchor
				case '#':
					val = string(inp[ist:i])
					kv[key] = val
					state = 15
					ist = i+1


				case '\r','\n':
					return res, fmt.Errorf("state 11 no key!")
				default:
			}


		// parse anchor
		case 15:
			switch inp[i] {

			case ' ':
				res.anchor = inp[ist:i]
				// parse protocol
				state = 3
				ist = i+1

			//error no protocol
			case '\r','\n':
				return res, fmt.Errorf("no protocol!")

			default:
			}

		default:
			return res, fmt.Errorf("invalid state %d", state)
		}
	}
	res.kv = kv
	return res, nil
}

func PrintRes(res resHttp) {
	fmt.Println("****** parse Http ******")
	fmt.Printf("cmd:     %s\n", res.cmd)
	fmt.Printf("domain:  %s\n", res.domain)
	fmt.Printf("path:    %s\n", res.path)
	fmt.Printf("anchor:  %s\n",res.anchor)
	fmt.Printf("proto:   %s\n", res.proto)
	fmt.Printf("http1.1: %t\n", res.prot1)
	fmt.Printf("  kv map\n")
	for k,v := range res.kv {
		fmt.Printf("  key: %s val: %s\n", k, v)
	}
	fmt.Println("**** end parse Http ****")
}
