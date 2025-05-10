package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "beginner"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("beginner")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("beginner") = %q, %v, want match for %q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Helllo("") = %q, %v, want '', error`, msg, err)
	}
}
