package color

import (
	"strings"
)

func AnsiiRGB(s string) ([]int, error) {
	color := ""
	end := len(s)
	if strings.Contains(s, `--color="`) {
		color = s[9:end]
	} else if strings.Contains(s, `--color=`) {
		color = s[8:end]
	}

	rgb, err := GetRGBValue(color)

	return rgb, err
}
