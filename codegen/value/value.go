package gen

import (
	"github.com/clipperhouse/typewriter"
)

func init() {
	templates = append(templates, &typewriter.Template{
		Name: "value",
		Text: `
func (self {{.Name}}) Equals(other Value) bool {
	if other, ok := other.({{.Name}}); ok {
		return self == other
	} else {
		return false
	}
}

func (v {{.Name}}) Ref() ref.Ref {
	return getRef(v)
}
`,
		TypeConstraint: typewriter.Constraint{Comparable: true},
	})
}
