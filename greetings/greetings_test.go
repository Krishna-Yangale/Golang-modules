package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Krishna"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Krishna")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want "",error`, msg, err)
	}
}
