package utils

import "iter"

func UNUSED(x ...any) {}

func Assert(condition bool) {
	if !condition {
		panic("Assertion failed")
	}
}

func Window[Slice ~[]E, E any](slice Slice, size int) iter.Seq[Slice] {
	return func(yield func(Slice) bool) {
		for i := range slice[:len(slice)-size+1] {
			if !yield(slice[i : i+size]) {
				return
			}
		}
	}
}
