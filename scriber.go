package scribe

import "io"

type Scriber interface {
	BeginDescribe(desc string)
	EndDescribe()
	Child(desc string) Scriber
	Print(str string)
	Printf(format string, args ...any)
	PrintLines(r io.Reader)
	Error(err error)
	Errorf(format string, args ...any)
}
