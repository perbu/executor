package executor_test

import (
	"testing"
	"github.com/perbu/executor"
)

func TestExecute(t *testing.T) {
	cases := []struct {
		name     string
		in       []byte
		wantErr  bool
		stdout   string
		stderr   string
		exitCode int
	}{
		{
			name:    `ok`,
			in:      []byte("#!/bin/sh\necho 'ok'"),
			wantErr: false,
		},
		{
			name:     `error, no output`,
			in:       []byte("#!/bin/sh\nexit 33"),
			wantErr:  true,
			stdout:   "",
			stderr:   "",
			exitCode: 33,
		},
		{
			name:     `error, with stderr output`,
			in:       []byte("#!/bin/sh\necho 'error' >&2\nexit 34"),
			wantErr:  true,
			stdout:   "",
			stderr:   "error\n",
			exitCode: 34,
		}, {
			name:     `error, with stdout output`,
			in:       []byte("#!/bin/sh\necho 'error'\nexit 35"),
			wantErr:  true,
			stdout:   "error\n",
			stderr:   "",
			exitCode: 35,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			got := executor.Execute(tt.in)
			if got != nil {
				// fmt.Println(fmt.Println(got.Error()))
				// Cast the error:
				e, ok := got.(*executor.ErrExecute)
				if !ok {
					t.Fatalf("Execute returned %T, want ErrExecute", got)
				}
				// check if the error is what it should be
				if e.Stdout != tt.stdout {
					t.Errorf("Execute(%s) = %v, want stdout: %s", tt.in, e.Stdout, tt.stdout)
				}
				if e.Stderr != tt.stderr {
					t.Errorf("Execute(%s) = %v, want stderr: %s", tt.in, e.Stderr, tt.stderr)
				}
				if e.ExitCode != tt.exitCode {
					t.Errorf("Execute(%s) = %v, want exit code: %d", tt.in, e.ExitCode, tt.exitCode)
				}

			}
			if tt.wantErr && got == nil {
				t.Errorf("Execute(%s) = %v, want error", tt.in, got)
			}

		})
	}
}

