package color

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestColor(t *testing.T) {
	tc := []struct {
		name string
		want string
		got  string
	}{
		{
			name: "BlackFg",
			want: "\033[1;30mfoo\033[0m",
			got:  BlackFg("foo"),
		},
		{
			name: "RedFg",
			want: "\033[1;31mfoo\033[0m",
			got:  RedFg("foo"),
		},
		{
			name: "GreenFg",
			want: "\033[1;32mfoo\033[0m",
			got:  GreenFg("foo"),
		},
		{
			name: "YellowFg",
			want: "\033[1;33mfoo\033[0m",
			got:  YellowFg("foo"),
		},
		{
			name: "BlueFg",
			want: "\033[1;34mfoo\033[0m",
			got:  BlueFg("foo"),
		},
		{
			name: "MagentaFg",
			want: "\033[1;35mfoo\033[0m",
			got:  MagentaFg("foo"),
		},
		{
			name: "CyanFg",
			want: "\033[1;36mfoo\033[0m",
			got:  CyanFg("foo"),
		},
		{
			name: "WhiteFg",
			want: "\033[1;37mfoo\033[0m",
			got:  WhiteFg("foo"),
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.got)
		})
	}
}
