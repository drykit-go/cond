package cond

// TODO: Make these funcs compatible with all types
// - via code generation
// - or wait for generics

// DefaultValue returns vzeroable if it's not a zero value, vdefault otherwise.
func DefaultValue(vzeroable, vdefault interface{}) interface{} {
	return Value(vzeroable, vdefault, !Zero(vzeroable))
}

// DefaultString returns vzeroable if it's not a zero value, vdefault otherwise.
func DefaultString(vzeroable, vdefault string) string {
	return String(vzeroable, vdefault, vzeroable != "")
}

// DefaultInt returns vzeroable if it's not a zero value, vdefault otherwise.
func DefaultInt(vzeroable, vdefault int) int {
	return Int(vzeroable, vdefault, vzeroable != 0)
}
