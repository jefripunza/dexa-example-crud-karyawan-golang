package helpers

func If[T any](cond bool, value_true, value_false T) T {
	if cond {
		return value_true
	}
	return value_false
}
