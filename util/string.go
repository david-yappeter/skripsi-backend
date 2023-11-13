package util

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func StringP(s string) *string {
	return &s
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}

func StringInSlice(s string, list []string) bool {
	for _, l := range list {
		if l == s {
			return true
		}
	}
	return false
}

func ParsePath(s string) string {
	return filepath.ToSlash(s)
}

func ParsePathTrim(s string) string {
	return strings.Trim(ParsePath(s), "/")
}

func ParsePathTrimLeft(s string) string {
	return strings.TrimLeft(ParsePath(s), "/")
}

func ParsePathTrimRight(s string) string {
	return strings.TrimRight(ParsePath(s), "/")
}

func RandomStringWithCharset(n int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func RandomString(n int) string {
	return RandomStringWithCharset(n, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
}

func CryptoRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}

		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func MustCryptoRandomString(n int) string {
	b, err := CryptoRandomString(n)
	if err != nil {
		panic(err)
	}

	return b
}

func NewSecretKey() string {
	return MustCryptoRandomString(64)
}

func WordWrap(sentence string, limit int) []string {
	// Check if the sentence is already within the character limit
	if utf8.RuneCountInString(sentence) <= limit {
		return []string{sentence}
	}

	var wrapped []string
	words := strings.Fields(sentence)
	currentLine := ""

	for _, word := range words {
		// Check if adding the current word exceeds the character limit
		if utf8.RuneCountInString(currentLine+" "+word) > limit {
			wrapped = append(wrapped, strings.TrimSpace(currentLine))
			currentLine = ""
		}

		// Append the word to the current line
		currentLine += " " + word
	}

	// Append the last line to the wrapped output
	if currentLine != "" {
		wrapped = append(wrapped, strings.TrimSpace(currentLine))
	}

	return wrapped
}

func LessText(texts []string, length int) string {
	textsLength := len(texts)
	if textsLength == 0 {
		return ""
	}

	if length <= 0 {
		length = textsLength
	}

	if length == 1 {
		lessText := texts[0]
		if textsLength > length {
			lessText += fmt.Sprintf(" and %d more", textsLength-length)
		}

		return lessText
	}

	minLength := length
	if textsLength < minLength {
		minLength = textsLength
	}

	lessText := texts[0]
	for i := 1; i < minLength-1; i++ {
		lessText += ", " + texts[i]
	}

	if textsLength > minLength {
		lessText += fmt.Sprintf(", %s and %d more", texts[minLength-1], textsLength-minLength)
	} else if minLength > 1 {
		lessText += fmt.Sprintf(" and %s", texts[minLength-1])
	}

	return lessText
}
