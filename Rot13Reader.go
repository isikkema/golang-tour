package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	nb := make([]byte, len(b))
	c := byte(0)
	capital := false
	n, err := r.r.Read(nb)
	for i := 0; i < n; i++ {
		capital = nb[i] >= 65 && nb[i] <= 90
		c = nb[i] + 13
		if (capital && c > 'Z') || (!capital && c > 'z') {
			c -= 26
		}

		b[i] = c
	}

	return n, err
}

func main() {
	fmt.Printf("%v %v\n", 'a', 'A')
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
