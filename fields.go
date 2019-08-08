package log

import "sort"

// Fields is log data
type Fields map[string]interface{}

// SortKeys returns an alphabetically sorted slice of keys
func (f Fields) SortKeys() []string {
	keys := make([]string, len(f))
	var i int
	for k := range f {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
