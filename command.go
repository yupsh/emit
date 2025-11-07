package command

import (
	"context"
	"io"

	yup "github.com/gloo-foo/framework"
)

type command struct {
	stdoutContent string
	stderrContent string
}

// Emit creates a command that produces output on both stdout and stderr.
// This is primarily useful for testing pipelines and demonstrating error handling.
//
// Example:
//
//	// Emit "hello" to stdout and "warning" to stderr
//	cmd := emit.Emit("hello", "warning")
//	yup.MustRun(cmd)
func Emit(stdoutContent, stderrContent string) yup.Command {
	return command{
		stdoutContent: stdoutContent,
		stderrContent: stderrContent,
	}
}

func (c command) Executor() yup.CommandExecutor {
	return func(ctx context.Context, stdin io.Reader, stdout, stderr io.Writer) error {
		// Write to stdout if content provided
		if c.stdoutContent != "" {
			if _, err := io.WriteString(stdout, c.stdoutContent); err != nil {
				return err
			}
			// Add newline if not already present
			if c.stdoutContent[len(c.stdoutContent)-1] != '\n' {
				if _, err := io.WriteString(stdout, "\n"); err != nil {
					return err
				}
			}
		}

		// Write to stderr if content provided
		if c.stderrContent != "" {
			if _, err := io.WriteString(stderr, c.stderrContent); err != nil {
				return err
			}
			// Add newline if not already present
			if c.stderrContent[len(c.stderrContent)-1] != '\n' {
				if _, err := io.WriteString(stderr, "\n"); err != nil {
					return err
				}
			}
		}

		return nil
	}
}

