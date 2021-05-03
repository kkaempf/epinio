// Package tracelog provides a logger for debugging and tracing
// This logger will not print anything, unless TRACE_LEVEL is at least 1
package tracelog

import (
	"context"
	"log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type CtxLoggerKey struct{}

// Logger returns the logger from the context, the server injects a logger into
// each request.
func Logger(ctx context.Context) logr.Logger {
	log, ok := ctx.Value(CtxLoggerKey{}).(logr.Logger)
	if !ok {
		// this should not happen, but let's be cautious
		return NewLogger().WithName("fallback")
	}
	return log
}

// TraceLevel returns the trace-level argument
func TraceLevel() int {
	return viper.GetInt("trace-level")
}

// LoggerFlags adds to viper flags
func LoggerFlags(pf *flag.FlagSet, argToEnv map[string]string) {
	pf.IntP("trace-level", "", 0, "Only print trace messages at or above this level (0 to 2, default 0)")
	viper.BindPFlag("trace-level", pf.Lookup("trace-level"))
	argToEnv["trace-level"] = "TRACE_LEVEL"
}

// NewServerLogger creates a new logger for server subcommand
func NewServerLogger() logr.Logger {
	return NewLogger().WithName("epinio")
}

// NewClientLogger creates a new logger with our setup
func NewClientLogger() logr.Logger {
	return NewLogger().WithName("EpinioClient")
}

// NewInstallClientLogger creates a new logger for the install subcommand
func NewInstallClientLogger() logr.Logger {
	return NewLogger().WithName("InstallClient")
}

// NewLogger creates a new logger with our setup
func NewLogger() logr.Logger {
	stdr.SetVerbosity(TraceLevel())
	return stdr.New(log.New(os.Stderr, "", log.LstdFlags)).V(1) // NOTE: Increment of level, not absolute.
}
