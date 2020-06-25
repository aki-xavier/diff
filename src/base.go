package diff

import "strings"

// Diff :
type Diff struct{}

// CreateDiff :
func CreateDiff() *Diff {
	d := &Diff{}
	return d
}

// Diff :
func (d *Diff) Diff(oldString, newString string) []*Change {
	ret := make([]*Change, 0)
	return ret
}

func (d *Diff) pushComponent(components []*Change, added, removed bool) {

}

func (d *Diff) extractCommon(basePath, newString, oldString, diagonalPath string) int {
	return 0
}

func (d *Diff) equals(left, right string) bool {
	return false
}

func (d *Diff) removeEmpty(array []interface{}) {

}

func (d *Diff) castInput(value string) string {
	return value
}

func (d *Diff) tokenize(value string) []string {
	return strings.Split(value, "")
}

func (d *Diff) join(chars []string) string {
	return strings.Join(chars, "")
}

func buildValues(diff *Diff, components []*Change, newString, oldString string, useLongestToken bool) {

}

func clonePath(path *Path) *Path {
	p := &Path{}
	return p
}
