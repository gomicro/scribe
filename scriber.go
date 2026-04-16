package scribe

import "io"

type Scriber interface {
	BeginDescribe(desc string)
	EndDescribe()
	Print(done string)
	PrintLines(r io.Reader)
	Error(err error)
}
