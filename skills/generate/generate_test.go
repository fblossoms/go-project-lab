package generate_test

import (
	"go18/skills/generate"
	"testing"
)

func TestStringSet(t *testing.T) {
	set := generate.NewSet[string]()
	set.Add("test")
	t.Log(set)
}

func TestIntSet(t *testing.T) {
	set := generate.NewSet[int]()
	set.Add(1)
	t.Log(set)
}
