package gocopy

import (
	"fmt"
	"io"
	"strings"
)

// ProgressCounter is struct for progress
type ProgressCounter struct {
	count uint64
}

// Write is implementation Writer interface for Progress
func (pc *ProgressCounter) Write(p []byte) (int, error) {
	n := len(p)
	pc.count += uint64(n)
	pc.PrintProgress()
	return n, nil
}

// PrintProgress is printing copy files
func (pc *ProgressCounter) PrintProgress() {
	fmt.Printf("\n complete... %d bytes", pc.count)
}

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

// Copier copies from rider to writer, returning the number of copied bytes and an error
func (gc *GoCopy) Copier() (int64, error) {

	if _, err := gc.R.Seek(gc.Offset, 0); err != nil {
		return 0, err
	}
	limitReader := io.LimitReader(gc.R, gc.Limit)
	counter := &ProgressCounter{}
	b := make([]byte, gc.Bs)
	written, err := io.CopyBuffer(gc.W, io.TeeReader(limitReader, counter), b)
	if err != nil {
		return 0, err
	}
	fmt.Printf("\n%s \n", strings.Repeat("-", 40))
	gc.W.Close()
	gc.R.Close()
	return written, nil

}
