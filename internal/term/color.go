package term

import (
	"github.com/fatih/color"
	"os"
)

type ForegroundColor int

const (
	FgBlack ForegroundColor = iota
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// printConfig and loggerOption work together to abstract away the library that
// we use to do things like colouring. By abstracting it, we can change the
// underlying library out later.
type printConfig struct {
	fgColor *ForegroundColor
}

type loggerOption func(config printConfig) printConfig

func WithForegroundColor(color ForegroundColor) loggerOption {
	return func(config printConfig) printConfig {
		config.fgColor = &color
		return config
	}
}

type Logger interface {
	// Print performs like any standard Print. Except, you can pass
	// LoggerOptions in with 'any' and they will be pulled out to be applied.
	//This allows you to use things like color.
	Print(a ...any) (n int, err error)

	// Printf performs like any standard Printf. Except, you can pass
	// LoggerOptions in with 'any' and they will be pulled out to be applied.
	//This allows you to use things like color.
	Printf(format string, a ...any) (n int, err error)

	//StartBuffer() Logger
	//FlushBuffer()
}

func NewStandardLogger() Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Print(a ...any) (n int, err error) {
	filteredAny, options := filterOutLoggerOption(a...)
	config := newPrintConfig(options...)
	c := newFatihColorPrinter(config)
	return c.Print(filteredAny...)
}

func (l *logger) Printf(format string, a ...any) (n int, err error) {
	filteredAny, options := filterOutLoggerOption(a...)
	config := newPrintConfig(options...)
	c := newFatihColorPrinter(config)
	return c.Fprintf(os.Stdout, format, filteredAny...)
}

func filterOutLoggerOption(a ...any) ([]any, []loggerOption) {
	var filteredAny []any
	var options []loggerOption
	for _, element := range a {
		if opt, ok := element.(loggerOption); ok {
			options = append(options, opt)
		} else {
			filteredAny = append(filteredAny, element)
		}
	}
	return filteredAny, options
}

func newPrintConfig(options ...loggerOption) printConfig {
	config := printConfig{}
	for _, opt := range options {
		config = opt(config)
	}
	return config
}

func newFatihColorPrinter(config printConfig) *color.Color {
	var attributes []color.Attribute

	if config.fgColor != nil {
		var fgColor color.Attribute
		switch *config.fgColor {
		case FgBlack:
			fgColor = color.FgBlack
		case FgRed:
			fgColor = color.FgRed
		case FgGreen:
			fgColor = color.FgGreen
		case FgYellow:
			fgColor = color.FgYellow
		case FgBlue:
			fgColor = color.FgBlue
		case FgMagenta:
			fgColor = color.FgMagenta
		case FgCyan:
			fgColor = color.FgCyan
		case FgWhite:
			fgColor = color.FgWhite
		}
		attributes = append(attributes, fgColor)
	}

	c := color.New(attributes...)

	return c
}
