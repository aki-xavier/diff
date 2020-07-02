package diff

// Change :
type Change struct {
	Value   string `json:"value"`
	Count   int    `json:"count"`
	Added   bool   `json:"added"`
	Removed bool   `json:"removed"`
}

func createChange() *Change {
	c := &Change{}
	c.Value = ""
	c.Count = 0
	c.Added = false
	c.Removed = false
	return c
}

func (c *Change) clone() *Change {
	cc := createChange()
	cc.Value = c.Value
	cc.Count = c.Count
	cc.Added = c.Added
	cc.Removed = c.Removed
	return cc
}
