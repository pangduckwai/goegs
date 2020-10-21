package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rdr rot13Reader) Read(b []byte) (n int, err error) {
	buf := make([]byte, len(b))

	cnt, err := rdr.r.Read(buf)
	if err == io.EOF {
		return 0, io.EOF
	}

	for i := 0; i < cnt; i++ {
		if buf[i] >= 65 && buf[i] <= 90 {
			b[i] = (buf[i]-52)%26 + 65
		} else if buf[i] >= 97 && buf[i] <= 122 {
			b[i] = (buf[i]-84)%26 + 97
		} else {
			b[i] = buf[i]
		}
	}
	return cnt, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
