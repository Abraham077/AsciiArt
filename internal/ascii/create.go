package ascii

import (
	"os"
	"strings"
)

func IsValidBanner(b string) bool {
	return b == "standard" || b == "shadow" || b == "thinkertoy"
}

func CreateAscii(text, banner string) (string, error) {
	text = strings.ReplaceAll(text, "\\n", "\n")
	for _, v := range text {
		if v == '\n' || v == '\r' {
			continue
		}
		if v < 32 || v > 126 {
			return "", os.ErrInvalid
		}
	}

	font, err := LoadBanner(banner)
	if err != nil {
		return "", err
	}
	height := CharHeight
	if glyph, ok := font[' ']; ok {
		height = len(glyph)
	}
	var b strings.Builder
	lines := strings.Split(text, "\n")

	for lin, line := range lines {
		if line == "" {
			if lin < len(lines)-1 {
				b.WriteRune('\n')
			}
			continue
		}
		rows := make([]string, height)

		for _, v := range line {
			glyph, ok := font[v]
			if !ok {
				continue
			}
			for i := 0; i < height && i < len(glyph); i++ {
				rows[i] += glyph[i]
			}
		}
		for i := 0; i < height; i++ {
			b.WriteString(rows[i])
			b.WriteRune('\n')
		}
	}
	result := b.String()
	result = strings.TrimRight(result, "\n")
	return result, nil

}
