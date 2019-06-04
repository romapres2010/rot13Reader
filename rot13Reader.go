package main

import (
	"io"
	"os"
	"strings"
	"time"
	"fmt"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read (b []byte) (len int, err error) {
	start := time.Now()
	
	len, err = r.r.Read(b)
	if err == io.EOF {
		return 0, err
	}
	
	t := time.Now()
	elapsed := t.Sub(start)
	
	fmt.Printf("len = %v, elapsed = %d ns\n", len, elapsed.Nanoseconds())
	
	return len, err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
