package cond

// TODO: Make these funcs compatible with all types
// - via code generation
// - or wait for generics

// Value returns vtrue if cond == true, vfalse otherwise.
func Value(vtrue, vfalse interface{}, cond bool) interface{} {
	if cond {
		return vtrue
	}
	return vfalse
}

// String returns vtrue if cond == true, vfalse otherwise.
func String(vtrue, vfalse string, cond bool) string {
	return Value(vtrue, vfalse, cond).(string)
}

// Int returns vtrue if cond == true, vfalse otherwise.
func Int(vtrue, vfalse int, cond bool) int {
	return Value(vtrue, vfalse, cond).(int)
}
