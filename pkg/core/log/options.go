package log

import (
	"io"
)

type Options struct {
	Level        Level
	Fields       map[string]interface{}
	Out          io.Writer
	ConsolePrint bool
}

type Option func(opts *Options)

func WithFields(fields map[string]interface{}) Option {
	return func(args *Options) {
		args.Fields = fields
	}
}

func WithLevel(level Level) Option {
	return func(args *Options) {
		args.Level = level
	}
}

func WithOutput(out io.Writer) Option {
	return func(args *Options) {
		args.Out = out
	}
}

func WithConsole(open bool) Option {
	return func(args *Options) {
		args.ConsolePrint = open
	}
}
