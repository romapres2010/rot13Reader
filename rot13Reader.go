package main

import (
	"io"
//	"os"
	"strings"
	"time"
	"fmt"
)

var lookup = map[byte]byte{
		'N': 'A', 'O': 'B', 'P': 'C', 'Q': 'D',
		'R': 'E', 'S': 'F', 'T': 'G', 'U': 'H',
		'V': 'I', 'W': 'J', 'X': 'K', 'Y': 'L',
		'Z': 'M', 'A': 'N', 'B': 'O', 'C': 'P',
		'D': 'Q', 'E': 'R', 'F': 'S', 'G': 'T',
		'H': 'U', 'I': 'V', 'J': 'W', 'K': 'X',
		'L': 'Y', 'M': 'Z', 'n': 'a', 'o': 'b',
		'p': 'c', 'q': 'd', 'r': 'e', 's': 'f',
		't': 'g', 'u': 'h', 'v': 'i', 'w': 'j',
		'x': 'k', 'y': 'l', 'z': 'm', 'a': 'n',
		'b': 'o', 'c': 'p', 'd': 'q', 'e': 'r',
		'f': 's', 'g': 't', 'h': 'u', 'i': 'v',
		'j': 'w', 'k': 'x', 'l': 'y', 'm': 'z',
	}

var enc = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var dec = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"

type rot13Reader struct {
	r io.Reader
}

func PrintElapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func (r *rot13Reader) Read (b []byte) (len int, err error) {
	
	//defer PrintElapsed(fmt.Sprintf("Read %v", r.r)) ()
	
	//start := time.Now()
	
	len, err = r.r.Read(b)
	
	if err == io.EOF {
		return 0, err
	}
	
	// map
//	for i, v := range b[:len] {
//		if rot13, ok := lookup[v]; ok {
//			b[i] = rot13
//		}
//	} 
	
	// IndexByte
	for i, v := range b[:len] {	
		if j := strings.IndexByte(enc, v); j >= 0 {
			b[i] = dec[j]
		}
	}
	

	//time.Sleep(time.Nanosecond  * 2)
	//end := time.Now()
	//elapsed := end.Sub(start)
	//fmt.Printf(" start = %v \n end   = %v \n len = %v, elapsed = %d ns\n", start, end, len, elapsed)
	
	return len, err
}

func test1 () {
	defer PrintElapsed("test1") ()
	
	for i := 0; i < 100000; i++ {
		s := strings.NewReader("Lbh penpxrq gur pbqr!")
		r := rot13Reader{s}
		
		b := make([]byte, 64)
		for {
			_, err := r.Read(b)
//			fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
//			fmt.Printf("b[:n] = %q\n", b[:n])
			if err == io.EOF {
				break
			}
		}
		
		//io.Copy(os.Stdout, &r)
	}
}

func main() {
	test1()
}
