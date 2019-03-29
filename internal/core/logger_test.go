package core

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	var (
		f = "%s-%s"
		v = []interface{}{"hello", "world"}
		l = NewLogger()

		iBuf bytes.Buffer
		wBuf bytes.Buffer
		eBuf bytes.Buffer
	)

	// capture output
	l.I.SetOutput(&iBuf)
	l.W.SetOutput(&wBuf)
	l.E.SetOutput(&eBuf)

	l.Info(v...)
	assert.Contains(t, iBuf.String(), "[INFO]")
	iBuf.Reset()

	l.Infof(f, v...)
	assert.Contains(t, iBuf.String(), "[INFO]")
	assert.Contains(t, iBuf.String(), "-")
	iBuf.Reset()

	l.Warn(v...)
	assert.Contains(t, wBuf.String(), "[WARN]")
	iBuf.Reset()

	l.Warnf(f, v...)
	assert.Contains(t, wBuf.String(), "[WARN]")
	assert.Contains(t, wBuf.String(), "-")
	iBuf.Reset()

	l.Error(v...)
	assert.Contains(t, eBuf.String(), "[ERROR]")
	iBuf.Reset()

	l.Errorf(f, v...)
	assert.Contains(t, eBuf.String(), "[ERROR]")
	assert.Contains(t, eBuf.String(), "-")
	iBuf.Reset()
}
