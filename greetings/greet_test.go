package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Yvonne"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Yvonne")
	if err == nil {
		fmt.Println(msg)
	}
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Yvonne") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
