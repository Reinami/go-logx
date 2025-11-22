package logx

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
)

type ColorLogger struct{}

func (c *ColorLogger) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (c *ColorLogger) Handle(_ context.Context, r slog.Record) error {
	ts := r.Time
	if ts.IsZero() {
		ts = time.Now()
	}

	timestamp := ts.Format("2006-01-02 15:04:05")
	color, levelLabel := levelColor(r.Level)

	fmt.Printf("%s %s[%s]%s %s\n",
		timestamp,
		color, levelLabel, colorReset,
		r.Message,
	)

	return nil
}

func (c *ColorLogger) WithAttrs(_ []slog.Attr) slog.Handler { return c }
func (c *ColorLogger) WithGroup(_ string) slog.Handler      { return c }

func levelColor(level slog.Level) (string, string) {
	switch {
	case level >= slog.LevelError:
		return colorRed, "ERROR"
	case level == slog.LevelWarn:
		return colorYellow, "WARN"
	case level == slog.LevelInfo:
		return colorGreen, "INFO"
	default:
		return colorBlue, "DEBUG"
	}
}
