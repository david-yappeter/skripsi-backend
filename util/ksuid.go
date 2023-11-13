package util

import "github.com/segmentio/ksuid"

// For filename
func NewKsuid() string {
	return ksuid.New().String()
}
