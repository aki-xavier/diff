package diff

import "fmt"

// Change :
type Change struct {
	Value   string
	Count   int
	Added   bool
	Removed bool
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

// String :
func (c *Change) String() string {
	return fmt.Sprintf(`{"value":"%v","count":%v,"added":%v,"removed":%v}`, c.Value, c.Count, c.Added, c.Removed)
}
