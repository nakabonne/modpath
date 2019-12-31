package modpath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name     string
		dir      string
		wantErr  bool
		expected string
	}{
		{
			name:     "general case",
			dir:      "./testdata/bar",
			wantErr:  false,
			expected: "example.com/foo/bar",
		},
		{
			name:    "wrong dir path",
			dir:     "./testdata/xxx",
			wantErr: true,
		},
		{
			name:     "no directory is explicitly given",
			dir:      "",
			wantErr:  false,
			expected: "github.com/nakabonne/modpath",
		},
		{
			name:    "wrong go.mod format",
			dir:     "./testdata/wrongfmt",
			wantErr: true,
		},
		{
			name:     "no go.mod file and then inspect parents",
			dir:      "./testdata/nomod",
			expected: "github.com/nakabonne/modpath",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m, err := Run(tc.dir)
			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expected, m)
		})
	}
}
