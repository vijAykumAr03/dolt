// Generated by: main
// TypeWriter: value
// Directive: +gen on Int32

package types

import (
	"github.com/attic-labs/noms/ref"
)

// DO NOT EDIT
//
// This file was generated by a tool.
// See http://clipperhouse.github.io/gen for details.

func (self Int32) Equals(other Value) bool {
	if other, ok := other.(Int32); ok {
		return self == other
	} else {
		return false
	}
}

func (v Int32) Ref() ref.Ref {
	return getRef(v)
}
