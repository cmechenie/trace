package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface capable of tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

//Off creates a tracer that ignores calls to Trace.
func Off() Tracer {
	return &nilTracer{}
}
