package utils

func BytesToStringPtr(bs []byte) *string {
	val := string(bs)
	return &val
}

func StringPtr(str string) *string {
	return &str
}

func StringPtrToVal(strPtr *string, def string) string {
	if strPtr == nil {
		return def
	}
	return *strPtr
}

func IntToInt32Ptr(i int) *int32 {
	val := int32(i)
	return &val
}
