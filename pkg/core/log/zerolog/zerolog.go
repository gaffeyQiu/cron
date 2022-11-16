package zerolog

import (
	"os"
	"reflect"
	"tsf-cron/pkg/core/helper"
	"tsf-cron/pkg/core/log"

	"github.com/rs/zerolog"
)

type ZeroLogger struct {
	opt log.Options
	l   zerolog.Logger
}

func New(opts ...log.Option) log.Logger {
	defaultOpt := log.Options{
		Level:        log.DebugLevel,
		Out:          os.Stdout,
		ConsolePrint: true,
	}

	for _, o := range opts {
		o(&defaultOpt)
	}

	//var out io.Writer

	if defaultOpt.ConsolePrint {
		if reflect.DeepEqual(defaultOpt.Out, os.Stdout) {
			defaultOpt.Out = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout})
		} else {
			defaultOpt.Out = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, defaultOpt.Out)
		}
	}

	zero := zerolog.New(defaultOpt.Out)

	for k, v := range defaultOpt.Fields {
		zero = zero.With().Interface(k, v).Logger()
	}
	//  4 是调试出来的
	zerolog.CallerSkipFrameCount = 4
	zero = zero.Level(transformZeroLogLevel(defaultOpt.Level)).
		With().
		Timestamp().
		Caller().
		Logger()

	return &ZeroLogger{
		opt: defaultOpt,
		l:   zero,
	}
}

func (z *ZeroLogger) GetOpt() log.Options {
	return z.opt
}

func (z *ZeroLogger) Fields(fields map[string]interface{}) log.Logger {
	for k, v := range fields {
		z.l = z.l.With().Interface(k, v).Logger()
	}
	return z
}

func (z *ZeroLogger) Debug(msg string, keyvals ...interface{}) {
	msgWithKvs(z.l.Debug(), msg, keyvals...)
}

func (z *ZeroLogger) Debugf(format string, val ...interface{}) {
	z.l.Debug().Msgf(format, val...)
}

func (z *ZeroLogger) Info(msg string, keyvals ...interface{}) {
	msgWithKvs(z.l.Info(), msg, keyvals...)
}

func (z *ZeroLogger) Infof(format string, val ...interface{}) {
	z.l.Info().Msgf(format, val...)
}

func (z *ZeroLogger) Warn(msg string, keyvals ...interface{}) {
	msgWithKvs(z.l.Warn(), msg, keyvals...)
}

func (z *ZeroLogger) Warnf(format string, val ...interface{}) {
	z.l.Warn().Msgf(format, val...)
}

func (z *ZeroLogger) Error(msg string, keyvals ...interface{}) {
	msgWithKvs(z.l.Error(), msg, keyvals...)
}

func (z *ZeroLogger) Errorf(format string, val ...interface{}) {
	z.l.Error().Msgf(format, val...)
}

func (z *ZeroLogger) Fatalf(format string, val ...interface{}) {
	z.l.Fatal().Msgf(format, val...)
}

func (z *ZeroLogger) String() string {
	return "zerolog"
}

func transformZeroLogLevel(l log.Level) zerolog.Level {
	switch l {
	case log.DebugLevel:
		return zerolog.DebugLevel
	case log.InfoLevel:
		return zerolog.InfoLevel
	case log.WarnLevel:
		return zerolog.WarnLevel
	case log.ErrorLevel:
		return zerolog.ErrorLevel
	case log.FatalLevel:
		return zerolog.FatalLevel
	}

	return zerolog.DebugLevel
}

func msgWithKvs(e *zerolog.Event, msg string, keyvals ...interface{}) {
	if len(keyvals)%2 == 0 {
		for i := 0; i < len(keyvals); i += 2 {
			e.Interface(helper.MustStringVar(keyvals[i]), keyvals[i+1])
		}
	}
	e.Msg(msg)
}
