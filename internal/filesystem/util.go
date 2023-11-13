package filesystem

import (
	"io"

	"github.com/gabriel-vasile/mimetype"
)

func guessMimeType(r io.ReadSeeker) (string, error) {
	// To guess the content type only the first 30720 bytes are used.
	buf := make([]byte, 30720)

	n, err := r.Read(buf)
	if err != nil {
		return "", err
	}
	r.Seek(0, io.SeekStart)

	mType := mimetype.Detect(buf[:n])

	return mType.String(), nil
}
