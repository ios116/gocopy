package copier

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// ReadCloseSeeker is interface ReadCloser, Seeker
type ReadCloseSeeker interface {
	io.ReadCloser
	io.Seeker
}

// GoCopy is structur
type GoCopy struct {
	R                 ReadCloseSeeker
	W                 io.WriteCloser
	Bs, Limit, Offset int64
}

// PrintProgress is printing copy files
func (gc *GoCopy) PrintProgress(written int64) {
	fmt.Printf("\n complete... %d bytes", written)
}

// Copier copies from rider to writer, returning the number of copied bytes and an error
func (gc *GoCopy) Copier() (written int64, err error) {

	if _, err := gc.R.Seek(gc.Offset, 0); err != nil {
		return 0, err
	}
	limitReader := io.LimitReader(gc.R, gc.Limit)

	b := make([]byte, gc.Bs)
	for {
		nr, er := limitReader.Read(b)

		if nr > 0 {
			nw, ew := gc.W.Write(b[0:nr])
			if nw > 0 {
				written += int64(nw)
				gc.PrintProgress(written)
			}
			if ew != nil {
				err = ew
				break
			}

			if nr != nw {
				err = errors.New("short write")
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	fmt.Printf("\n%s\n", strings.Repeat("-", 40))
	gc.W.Close()
	gc.R.Close()
	return written, nil

}
