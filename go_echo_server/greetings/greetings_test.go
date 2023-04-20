package greetings

import (
	"regexp"
	"testing"
)

const (
	starting = "Starting up..."
)

func TestHelloTrue(t *testing.T) {
	want := regexp.MustCompile(starting)
	msg, err := Hello(true)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(true) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloFalse(t *testing.T) {
	msg, err := Hello(false)
	if msg != "" || err == nil {
		t.Fatalf(`Hello(false) = %q, %v, want "", error`, msg, err)
	}
}
