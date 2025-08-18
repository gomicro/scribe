package scribe

type Theme struct {
	Describe func(string) string
	Print    func(string) string
	Error    func(error) string
}

var DefaultTheme = &Theme{
	Describe: NoopDecorator,
	Print:    NoopDecorator,
	Error:    NoopErrDecorator,
}

func NoopDecorator(s string) string {
	return s
}

func NoopErrDecorator(err error) string {
	return err.Error()
}
