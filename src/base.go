package diff

import (
	"strings"
	"unicode/utf8"
)

// BaseCustomize :
type BaseCustomize struct{}

func (c *BaseCustomize) tokenize(value string) []string {
	return strings.Split(value, "")
}
func (c *BaseCustomize) join(array []string) string {
	return strings.Join(array, "")
}
func (c *BaseCustomize) castInput(value string) string {
	return value
}
func (c *BaseCustomize) equals(left string, right string) bool {
	return left == right
}

// Diff :
type Diff struct {
	customize       Customize
	oldStrings      []string
	newStrings      []string
	oldLen          int
	newLen          int
	editLength      int
	bestPathes      []*Path
	useLongestToken bool
}

// CreateDiff :
func CreateDiff() *Diff {
	d := &Diff{}
	d.customize = &BaseCustomize{}
	d.oldStrings = make([]string, 0)
	d.newStrings = make([]string, 0)
	d.oldLen = 0
	d.newLen = 0
	d.editLength = 1
	d.bestPathes = make([]*Path, 0)
	d.useLongestToken = false
	return d
}

// Diff :
func (d *Diff) Diff(oldString, newString string) []*Change {
	oldString = d.customize.castInput(oldString)
	newString = d.customize.castInput(newString)
	d.oldStrings = d.removeEmpty(d.customize.tokenize(oldString))
	d.newStrings = d.removeEmpty(d.customize.tokenize(newString))
	d.newLen = len(d.newStrings)
	d.oldLen = len(d.oldStrings)
	d.editLength = 1
	maxEditLength := d.newLen + d.oldLen

	d.bestPathes = d.bestPathes[:0]
	d.bestPathes = append(d.bestPathes, CreatePath())
	ret := make([]*Change, 0)
	oldPos := d.extractCommon(d.bestPathes[0], d.newStrings, d.oldStrings, 0)
	if d.bestPathes[0].newPos+1 >= d.newLen && oldPos+1 >= d.oldLen {
		c := createChange()
		c.Value = d.customize.join(d.newStrings)
		c.Count = d.newLen
		ret = append(ret, c)
		return ret
	}

	for d.editLength <= maxEditLength {
		_ret := d.execEditLength()
		if _ret != nil {
			return _ret
		}
	}
	return ret
}

func (d *Diff) execEditLength() []*Change {
	for diagonalPath := -1 * d.editLength; diagonalPath <= d.editLength; diagonalPath += 2 {
		var basePath *Path
		var addPath *Path
		var removePath *Path
		if diagonalPath-1 >= 0 && diagonalPath-1 <= len(d.bestPathes)-1 {
			addPath = d.bestPathes[diagonalPath-1]
		}
		if diagonalPath+1 >= 0 && diagonalPath+1 <= len(d.bestPathes)-1 {
			removePath = d.bestPathes[diagonalPath+1]
		}
		var oldPos int
		if removePath != nil {
			oldPos = removePath.newPos - diagonalPath
		} else {
			oldPos = 0 - diagonalPath
		}

		if addPath != nil {
			d.bestPathes[diagonalPath-1] = nil
		}

		var canAdd bool
		var canRemove bool
		if addPath != nil && addPath.newPos+1 < d.newLen {
			canAdd = true
		} else {
			canAdd = false
		}
		if removePath != nil && 0 <= oldPos && oldPos < d.oldLen {
			canRemove = true
		} else {
			canRemove = false
		}
		if !canAdd && !canRemove {
			if diagonalPath >= 0 && diagonalPath <= len(d.bestPathes)-1 {
				d.bestPathes[diagonalPath] = nil
			}
			continue
		}

		if !canAdd || canRemove && addPath.newPos < removePath.newPos {
			basePath = clonePath(removePath)
			basePath.pushComponent(false, true)
		} else {
			basePath = addPath
			basePath.newPos++
			basePath.pushComponent(true, false)
		}

		oldPos = d.extractCommon(basePath, d.newStrings, d.oldStrings, diagonalPath)

		if basePath.newPos+1 >= d.newLen && oldPos+1 >= d.oldLen {
			return buildValues(d, basePath.components, d.newStrings, d.oldStrings, d.useLongestToken)
		}
		if diagonalPath >= 0 && diagonalPath <= len(d.bestPathes)-1 {
			d.bestPathes[diagonalPath] = basePath
		} else if diagonalPath == len(d.bestPathes) {
			d.bestPathes = append(d.bestPathes, basePath)
		}
	}
	d.editLength++
	return nil
}

func (d *Diff) extractCommon(basePath *Path, newStrings, oldStrings []string, diagonalPath int) int {
	newLen := len(newStrings)
	oldLen := len(oldStrings)
	newPos := basePath.newPos
	oldPos := newPos - diagonalPath
	commonCount := 0

	for newPos+1 < newLen && oldPos+1 < oldLen && d.customize.equals(newStrings[newPos+1], oldStrings[oldPos+1]) {
		newPos++
		oldPos++
		commonCount++
	}

	if commonCount != 0 {
		change := createChange()
		change.Count = commonCount
		basePath.components = append(basePath.components, change)
	}

	basePath.newPos = newPos
	return oldPos
}

func (d *Diff) removeEmpty(array []string) []string {
	ret := make([]string, 0)
	for _, arr := range array {
		if arr != "" {
			ret = append(ret, arr)
		}
	}
	return ret
}

func buildValues(diff *Diff, components []*Change, newStrings, oldStrings []string, useLongestToken bool) []*Change {
	componentPos := 0
	componentLen := len(components)
	newPos := 0
	oldPos := 0

	for ; componentPos < componentLen; componentPos++ {
		component := components[componentPos]
		if !component.Removed {
			if !component.Added && useLongestToken {
				value := newStrings[newPos : newPos+component.Count]
				for i, v := range value {
					oldValue := oldStrings[oldPos+i]
					if utf8.RuneCountInString(oldValue) > utf8.RuneCountInString(v) {
						value[i] = oldValue
					}
				}
				component.Value = diff.customize.join(value)
			} else {
				component.Value = diff.customize.join(newStrings[newPos : newPos+component.Count])
			}
			newPos += component.Count

			if !component.Added {
				oldPos += component.Count
			}
		} else {
			component.Value = diff.customize.join(oldStrings[oldPos : oldPos+component.Count])
			oldPos += component.Count

			if componentPos != 0 && components[componentPos-1].Added {
				tmp := components[componentPos-1]
				components[componentPos-1] = components[componentPos]
				components[componentPos] = tmp
			}
		}
	}

	lastComponent := components[componentLen-1]
	if componentLen > 1 && (lastComponent.Added || lastComponent.Removed) && diff.customize.equals("", lastComponent.Value) {
		components[componentLen-2].Value += lastComponent.Value
		components = components[:len(components)-1]
	}

	return components
}

func clonePath(path *Path) *Path {
	p := &Path{}
	p.newPos = path.newPos
	for _, change := range path.components {
		p.components = append(p.components, change.clone())
	}
	return p
}
