package ascii

import (
	"os"
	"strings"
)

const CharHeight = 8

func LoadBanner(name string) (map[rune][]string, error) {
	filename := "banners/" + name + ".txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
	charMap := make(map[rune][]string)

	for i := 0; i < 95; i++ {
		ch := rune(32 + i)
		start := 1 + i*(CharHeight+1)
		end := start + CharHeight
		if end > len(lines) {
			break
		}
		charMap[ch] = lines[start:end]
	}

	return charMap, nil
}
