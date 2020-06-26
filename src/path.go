package diff

// Path :
type Path struct {
	newPos     int
	components []*Change
}

// CreatePath :
func CreatePath() *Path {
	p := &Path{}
	p.newPos = -1
	p.components = make([]*Change, 0)
	return p
}

func (p *Path) pushComponent(added, removed bool) {
	var last *Change
	if len(p.components) != 0 {
		last = p.components[len(p.components)-1]
		if last.Added == added && last.Removed == removed {
			p.components[len(p.components)-1].Count++
		} else {
			last = nil
		}
	}
	if last == nil {
		change := createChange()
		change.Count = 1
		change.Added = added
		change.Removed = removed
		p.components = append(p.components, change)
	}
}
