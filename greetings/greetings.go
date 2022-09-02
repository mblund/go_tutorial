package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"rsc.io/quote"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	//Return a greeting that embeds the name in a message.
	//var message string

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	message := quote.Go()

	formats := []string{
		"Hi, %v. Welcome! " + message,
		"Great to see you, %v! " + message,
		"Hail, %v! Well met! " + message,
	}

	return formats[rand.Intn(len(formats))]
}
