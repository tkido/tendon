package ui

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"strconv"
)

var (
	colorCache    map[string]color.Color
	colorCodeRune map[rune]struct{}
)

func init() {
	colorCache = map[string]color.Color{}
	colorCodeRune = map[rune]struct{}{}
	for _, r := range "0123456789abcdefABCDEF" {
		colorCodeRune[r] = struct{}{}
	}
}

// Color return color
func Color(s string) color.Color {
	raw := s
	if c, ok := colorCache[s]; ok {
		return c
	}
	for _, r := range s {
		if _, ok := colorCodeRune[r]; !ok {
			log.Printf("ui.Color: invalid rune %s for color code", string(r))
			return color.Black
		}
	}
	switch len(s) {
	case 3, 4:
		rs := []rune{}
		for _, r := range s {
			rs = append(rs, r, r)
		}
		s = string(rs)
	case 6, 8:
	default:
		log.Printf("ui.Color: invalid color code %s", s)
		return color.Black
	}
	us := [4]uint8{}
	for i := 0; i < len(s)/2; i++ {
		ui, _ := strconv.ParseUint(s[i*2:i*2+2], 16, 8)
		us[i] = uint8(ui)
	}
	if len(s) == 6 {
		us[3] = 0xff
	}
	c := color.NRGBA{us[0], us[1], us[2], us[3]}
	colorCache[raw] = c
	return c
}

// ColorCode show RGBA color as HTML color code.
// e.g. White -> #ffffffff
func ColorCode(c color.Color) string {
	if c == nil {
		return "#nil"
	}
	rgba := [4]uint32{}
	rgba[0], rgba[1], rgba[2], rgba[3] = c.RGBA()
	buf := bytes.Buffer{}
	buf.WriteString("#")
	for _, u := range rgba {
		buf.WriteString(fmt.Sprintf("%02x", uint8(u)))
	}
	return buf.String()
}
