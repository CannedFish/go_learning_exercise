package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	rb := make([]byte, len(b))
	n, err := rot.r.Read(rb)
	if err == io.EOF {
		return 0, io.EOF
	}
	for i := 0; i < n; i++ {
		if rb[i] >= 65 && rb[i] <= 90 {
			b[i] = (rb[i] - 65 + 13) % 26 + 65
		} else if rb[i] >= 97 && rb[i] <= 122 {
			b[i] = (rb[i] - 97 + 13) % 26 + 97
		} else {
			b[i] = rb[i]
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}