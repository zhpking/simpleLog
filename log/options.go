package log

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

const FmtEmptySeparate = ""

type Level uint8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var LevelNameMapping = map[Level]string {
	DebugLevel:"DEBUG",
	InfoLevel:"INFO",
	WarnLevel:"WARN",
	ErrorLevel:"ERROR",
	PanicLevel:"PANIC",
	FatalLevel:"FATAL",
}

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO":
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}

	return true
}

func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}

	return nil
}

// type Formatter func(params LogFormatterParams) string
// type Formatter func() string

type options struct {
	output io.Writer
	level Level
	stdLevel Level
	formatter Formatter
	disableCaller bool
}

type Option func(*options)

func initOption(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}

/*
设置输出级别
*/
func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

/*
设置输出位置
*/
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

/*
设置输出格式
*/
func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

/*
设置是否打印文件名和行号
*/
func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}
