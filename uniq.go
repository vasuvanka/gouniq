package gouniq

import (
	"crypto/rand"
	"errors"
)

var chars = []byte("_-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// NewID returns uniq string ID consisting of standard characters.
func NewID(length int) (string, error) {
	if length == 0 {
		return "", errors.New("empty string")
	}
	var id string
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	for i := length - 1; i >= 0; i-- {
		id += string(chars[bytes[i]&63])
	}
	return id, nil
}
