package diff

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
