package clog

import (
	"fmt"
	"go-fiber-ent-web-layout/internal/tools/pool"
	"io"
	"log/slog"
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
		pool.Go(func() {
			func(write io.Writer) {
				if _, err := write.Write(p); err != nil {
					slog.Error(fmt.Sprintf("logger write errer message:%s", err))
				}
			}(wr)
		})
	}
	return os.Stdout.Write(p)
}
