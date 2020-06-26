package diff

import (
	"strings"
)

// LineCustomize :
type LineCustomize struct {
	BaseCustomize
}

func (l *LineCustomize) tokenize(value string) []string {
	linesAndNewLines := strings.Split(value, "\n")

	if len(linesAndNewLines) != 0 && linesAndNewLines[len(linesAndNewLines)-1] == "" {
		linesAndNewLines = linesAndNewLines[:len(linesAndNewLines)-1]
	}
	return linesAndNewLines
}

func (l *LineCustomize) join(array []string) string {
	return strings.Join(array, "\n")
}

// CompareLines :
func CompareLines(oldStr, newStr string) []*Change {
	d := CreateDiff()
	d.customize = &LineCustomize{}
	return d.Diff(oldStr, newStr)
}
