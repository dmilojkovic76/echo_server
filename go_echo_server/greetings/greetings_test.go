package greetings

import (
	"regexp"
	"testing"
)

const (
	starting_msg = "Starting up..."
)

func TestHelloTrue(t *testing.T) {
	want := regexp.MustCompile(starting_msg)
	msg := StartUp()
	if !want.MatchString(msg) {
		t.Fatalf(`Hello() = %q, %v want match for "%q"`, msg, want, starting_msg)
	}
}
