package diff

import "strings"

// ArrayCustomize :
type ArrayCustomize struct {
	BaseCustomize
}

func (a *ArrayCustomize) tokenize(value string) []string {
	return strings.Split(value, "\t")
}

func (a *ArrayCustomize) join(array []string) string {
	return strings.Join(array, "\t")
}

// CompareArray :
func CompareArray(oldArr, newArr []string) []*Change {
	d := CreateDiff()
	d.Customize = &ArrayCustomize{}
	return d.Diff(strings.Join(oldArr, "\t"), strings.Join(newArr, "\t"))
}
