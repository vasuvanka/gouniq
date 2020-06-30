package gouniq

import (
	"crypto/rand"
	"errors"
)

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// NewID returns uniq string ID consisting of standard characters.
func NewID(length int) (string, error) {
	return newLenChars(length)
}

func newLenChars(length int) (string, error) {
	if length == 0 {
		return "", errors.New("empty string")
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		return "", errors.New("invalid length")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return "", err
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b), nil
			}
		}
	}
}
