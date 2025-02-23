package scribe

import (
	"fmt"
	"io"
	"strings"
)

type Scribe struct {
	writer io.Writer
	level  int
}

func NewScribe(writer io.Writer) Scriber {
	return &Scribe{
		writer: writer,
	}
}

func (s *Scribe) BeginDescribe(desc string) {
	s.println()
	s.printt(desc)
	s.level++
}

func (s *Scribe) EndDescribe() {
	s.level--
}

func (s *Scribe) Done(done string) {
	s.level++
	s.printt(done)
	s.level--
}

func (s *Scribe) print(str string) {
	fmt.Fprintf(s.writer, "%v\n", str)
}

func (s *Scribe) printt(str string) {
	s.print(fmt.Sprintf("%v%v", s.space(), str))
}

func (s *Scribe) println() {
	s.print("")
}

func (s *Scribe) space() string {
	return strings.Repeat(" ", s.level*2)
}
