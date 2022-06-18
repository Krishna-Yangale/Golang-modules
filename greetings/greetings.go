package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty name:")
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Multi(names []string) (map[string]string, error) {

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

func randomFormat() string {
	formats := []string{
		"good morning %v",
		"good evening %v",
		"Have a nice day %v",
	}
	return formats[rand.Intn(len(formats))]
}
