package greetings

import (
	"errors"
)

func Hello(b bool) (string, error) {
	// A dummy check to satisfy return signature with an error
	if !b {
		return "", errors.New("There was an error!")
	}

	// Return a message and no error
	return "Starting up...", nil
}
