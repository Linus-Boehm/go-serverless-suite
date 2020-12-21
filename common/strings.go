package common

func StringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func StringPtr(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}
