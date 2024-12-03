package utils

func UNUSED(x ...any) {}

func Assert(condition bool) {
	if !condition {
		panic("Assertion failed")
	}
}
