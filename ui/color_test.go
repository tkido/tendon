package ui

import (
	"image/color"
	"testing"
)

func TestColor(t *testing.T) {
	cases := []struct {
		Put  string
		Want color.Color
	}{
		{"fff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"ffff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"ffffff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"ffffffff", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"fff0", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"ffffff00", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"FFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFFFFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFFFFFFF", color.NRGBA{0xff, 0xff, 0xff, 0xff}},
		{"FFF0", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"FFFFFF00", color.NRGBA{0xff, 0xff, 0xff, 0x00}},
		{"abcdef", color.NRGBA{0xab, 0xcd, 0xef, 0xff}},
		{"ABCDEF", color.NRGBA{0xAB, 0xCD, 0xEF, 0xFF}},
	}
	for _, c := range cases {
		put := c.Put
		got := Color(c.Put)
		want := c.Want
		if got != want {
			t.Errorf("put %v got %v want %v", put, got, want)
		}
	}
}

func TestColorCode(t *testing.T) {
	cases := []struct {
		Put  color.Color
		Want string
	}{
		{color.Black, "#000000ff"},
		{color.White, "#ffffffff"},
		{color.NRGBA{0xff, 0x00, 0x00, 0xff}, "#ff0000ff"},
		{color.NRGBA{0x00, 0xff, 0x00, 0xff}, "#00ff00ff"},
		{color.NRGBA{0x00, 0x00, 0xff, 0xff}, "#0000ffff"},
		{color.NRGBA{0x00, 0x00, 0x00, 0x00}, "#00000000"},
		{color.NRGBA{0xff, 0xff, 0xff, 0xff}, "#ffffffff"},
	}
	for _, c := range cases {
		put := c.Put
		got := ColorCode(c.Put)
		want := c.Want
		if got != want {
			t.Errorf("put %v got %v want %v", put, got, want)
		}
	}
}
