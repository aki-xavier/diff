package diff

// Customize :
type Customize interface {
	tokenize(string) []string
	join([]string) string
	castInput(string) string
	equals(string, string) bool
}
