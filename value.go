package wrike

func Bool(v bool) *bool { return &v }
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

func String(v string) *string { return &v }
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func Int(v int) *int { return &v }
func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}
