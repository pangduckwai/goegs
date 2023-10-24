package common

import "fmt"

// Err config errors
type Err struct {
	Hdr   string
	Msg   []string
	Fatal bool // Severity level: true -> fatal error, exit immediately
}

func (e *Err) Error() string {
	s := ""
	if e.Fatal {
		s = "[Fatal]"
	}

	if len(e.Msg) == 1 {
		return fmt.Sprintf("[%v]%v %v", e.Hdr, s, e.Msg[0])
	}

	msg := fmt.Sprintf("[%v]%v", e.Hdr, s)
	for _, m := range e.Msg {
		msg = fmt.Sprintf("%v\n%v", msg, m)
	}
	return msg
}

func NewError(fatal bool, header string, message ...string) (err *Err) {
	return &Err{
		header, message, fatal,
	}
}
