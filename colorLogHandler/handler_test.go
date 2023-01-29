package colorLogHandler

import (
	"bytes"
	"errors"
	"github.com/klusoga-software/klusoga-logger/colors"
	"golang.org/x/exp/slog"
	"testing"
)

type testWriter struct {
	Buffer bytes.Buffer
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	if t.Buffer.Len() > 0 {
		t.Buffer.Reset()
	}
	return t.Buffer.Write(p)
}

func TestHandler(t *testing.T) {
	writer := testWriter{Buffer: bytes.Buffer{}}

	slog.SetDefault(slog.New(NewColorLogHandler(&writer)))

	slog.Error("An Error occurred", errors.New("test"))

	if writer.Buffer.String() != colors.Red+" Level=\"ERROR\" Message=\"An Error occurred\" err=\"test\" "+colors.Clear {
		t.Log(writer.Buffer.String())
		t.Fail()
	}

	slog.Info("This is an Info")

	if writer.Buffer.String() != colors.Green+" Level=\"INFO\" Message=\"This is an Info\"  "+colors.Clear {
		t.Log(writer.Buffer.String())
		t.Fail()
	}

	slog.Debug("This is a debug message", "os", "mac")

	if writer.Buffer.String() != colors.Blue+" Level=\"DEBUG\" Message=\"This is a debug message\" os=\"mac\" "+colors.Clear {
		t.Log(writer.Buffer.String())
		t.Fail()
	}

	slog.Warn("This is a warning")

	if writer.Buffer.String() != colors.Yellow+" Level=\"WARN\" Message=\"This is a warning\"  "+colors.Clear {
		t.Log(writer.Buffer.String())
		t.Fail()
	}
}
