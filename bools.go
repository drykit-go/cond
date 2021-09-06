package cond

// XOR returns true if a != b
func XOR(a, b bool) bool {
	return a != b
}

// And returns true if all bools are true, false otherwise.
func And(bools ...bool) bool {
	return !NotAll(bools...)
}

// None returns true if all bools are false, false otherwise.
func None(bools ...bool) bool {
	return !Or(bools...)
}

// NotAll returns true if any bool is false, false otherwise.
func NotAll(bools ...bool) bool {
	return untilBool(func(b bool) bool { return !b }, bools...)
}

// Or returns true if any bool is true, false otherwise.
func Or(bools ...bool) bool {
	return untilBool(func(b bool) bool { return b }, bools...)
}

func untilBool(stop func(bool) bool, bools ...bool) bool {
	for _, b := range bools {
		if stop(b) {
			return true
		}
	}
	return false
}
