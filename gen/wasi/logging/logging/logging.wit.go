// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package logging represents the imported interface "wasi:logging/logging".
//
// WASI Logging is a logging API intended to let users emit log messages with
// simple priority levels and context values.
package logging

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// Level represents the enum "wasi:logging/logging#level".
//
// A log level, describing a kind of message.
//
//	enum level {
//		trace,
//		debug,
//		info,
//		warn,
//		error,
//		critical
//	}
type Level uint8

const (
	// Describes messages about the values of variables and the flow of
	// control within a program.
	LevelTrace Level = iota

	// Describes messages likely to be of interest to someone debugging a
	// program.
	LevelDebug

	// Describes messages likely to be of interest to someone monitoring a
	// program.
	LevelInfo

	// Describes messages indicating hazardous situations.
	LevelWarn

	// Describes messages indicating serious errors.
	LevelError

	// Describes messages indicating fatal errors.
	LevelCritical
)

var stringsLevel = [6]string{
	"trace",
	"debug",
	"info",
	"warn",
	"error",
	"critical",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e Level) String() string {
	return stringsLevel[e]
}

// Log represents the imported function "log".
//
// Emit a log message.
//
// A log message has a `level` describing what kind of message is being
// sent, a context, which is an uninterpreted string meant to help
// consumers group similar messages, and a string containing the message
// text.
//
//	log: func(level: level, context: string, message: string)
//
//go:nosplit
func Log(level Level, context string, message string) {
	level0 := (uint32)(level)
	context0, context1 := cm.LowerString(context)
	message0, message1 := cm.LowerString(message)
	wasmimport_Log((uint32)(level0), (*uint8)(context0), (uint32)(context1), (*uint8)(message0), (uint32)(message1))
	return
}

//go:wasmimport wasi:logging/logging log
//go:noescape
func wasmimport_Log(level0 uint32, context0 *uint8, context1 uint32, message0 *uint8, message1 uint32)
