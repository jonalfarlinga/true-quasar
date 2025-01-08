package common

import (
	"strings"
)

type Interactable interface {
	GetBounds() (float32, float32, float32, float32)
}

func Collide(x, y int, button Interactable) bool {
	left, top, right, bottom := button.GetBounds()
	X, Y := float32(x), float32(y)
	return X > left && X < right && Y > top && Y < bottom
}

func WrapText(text string, lineLen int) string {
	var result string
	var line string
	words := strings.Fields(text)
	for _, word := range words {
		if len(line)+len(word)+1 > lineLen {
			result += line + "\n"
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		result += line
	}
	return result
}

func WrapLines(text string, lineLen int) []string {
	var result []string
	var line string
	words := strings.Fields(text)
	for _, word := range words {
		if len(line)+len(word)+1 > lineLen {
			result = append(result, line)
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		result = append(result, line)
	}
	return result
}
