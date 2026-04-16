package color

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestColor(t *testing.T) {
	t.Parallel()

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
		{
			name: "HiBlackFg",
			want: "\033[1;90mfoo\033[0m",
			got:  HiBlackFg("foo"),
		},
		{
			name: "HiRedFg",
			want: "\033[1;91mfoo\033[0m",
			got:  HiRedFg("foo"),
		},
		{
			name: "HiGreenFg",
			want: "\033[1;92mfoo\033[0m",
			got:  HiGreenFg("foo"),
		},
		{
			name: "HiYellowFg",
			want: "\033[1;93mfoo\033[0m",
			got:  HiYellowFg("foo"),
		},
		{
			name: "HiBlueFg",
			want: "\033[1;94mfoo\033[0m",
			got:  HiBlueFg("foo"),
		},
		{
			name: "HiMagentaFg",
			want: "\033[1;95mfoo\033[0m",
			got:  HiMagentaFg("foo"),
		},
		{
			name: "HiCyanFg",
			want: "\033[1;96mfoo\033[0m",
			got:  HiCyanFg("foo"),
		},
		{
			name: "HiWhiteFg",
			want: "\033[1;97mfoo\033[0m",
			got:  HiWhiteFg("foo"),
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.got)
		})
	}
}
