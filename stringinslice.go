package stringinslice

//StringInSlice checks a slice of string if string in array1 already part of new slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
