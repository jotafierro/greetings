package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Jota"
	want := regexp.MustCompile(`\b` + name + `\b`) // \b is a word boundary - busca por una plabra completa en este caso el nombre

	msg, err := Hello("Jota")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Jota") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
