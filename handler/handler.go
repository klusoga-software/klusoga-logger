package handler

import (
	"context"
	"fmt"
	"github.com/klusoga-software/klusoga-logger/colors"
	"golang.org/x/exp/slog"
	"io"
)

type colorLogHandler struct {
	writer io.Writer
}

func NewColorLogHandler(writer io.Writer) *colorLogHandler {
	return &colorLogHandler{writer: writer}
}

func (c *colorLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (c *colorLogHandler) Handle(r slog.Record) error {
	var colorPrefix string

	switch r.Level {
	case slog.LevelInfo:
		colorPrefix = colors.Green
		break
	case slog.LevelError:
		colorPrefix = colors.Red
		break
	case slog.LevelWarn:
		colorPrefix = colors.Yellow
		break
	case slog.LevelDebug:
		colorPrefix = colors.Blue
		break
	}

	attributes := ""

	r.Attrs(func(attr slog.Attr) {
		attributes += attr.Key + "="
		attributes += "\"" + attr.Value.String() + "\""
	})

	_, err := c.writer.Write([]byte(fmt.Sprintf("%s Level=\"%s\" Message=\"%s\" %s %s", colorPrefix, r.Level.String(), r.Message, attributes, colors.Clear)))
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *colorLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return c
}

func (c *colorLogHandler) WithGroup(name string) slog.Handler {
	return c
}
