package clog

import (
	"io"
	"os"
)

type CustomSlogWriter struct {
	writes []io.Writer
}

func (w *CustomSlogWriter) RegisterWriter(write io.Writer) *CustomSlogWriter {
	w.writes = append(w.writes, write)
	return w
}

func (w *CustomSlogWriter) Write(p []byte) (int, error) {
	for _, wr := range w.writes {
		go func(write io.Writer) {
			_, _ = write.Write(p)
		}(wr)
	}
	return os.Stdout.Write(p)
}
