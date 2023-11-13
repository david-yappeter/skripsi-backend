package util

import "github.com/oklog/ulid/v2"

// For filename
func NewUlid() string {
	return ulid.Make().String()
}
