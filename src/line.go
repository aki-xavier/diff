package diff

import (
	"strings"
)

// LineCustomize :
type LineCustomize struct {
	BaseCustomize
}

func (l *LineCustomize) tokenize(value string) []string {
	retLines := make([]string, 0)
	values := strings.Split(value, "\n")
	for i, v := range values {
		if i == len(values)-1 {
			if v != "" {
				retLines = append(retLines, v)
			}
		} else {
			retLines = append(retLines, v+"\n")
		}
	}
	return retLines
}

// CompareLines :
func CompareLines(oldStr, newStr string) []*Change {
	d := CreateDiff()
	d.Customize = &LineCustomize{}
	return d.Diff(oldStr, newStr)
}
