package util

import "io"

// ReadSeekNopCloser returns a ReadSeekCloser with a no-op Close method wrapping
// the provided ReadSeeker r.
func ReadSeekNopCloser(r io.ReadSeeker) io.ReadSeekCloser {
	return readSeekNopCloser{r}
}

type readSeekNopCloser struct {
	io.ReadSeeker
}

func (readSeekNopCloser) Close() error { return nil }
