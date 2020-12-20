package whicharein

import (
	"sort"
	"stringinslice"
	"strings"
)

//InArray
func InArray(array1 []string, array2 []string) []string {
	newString := []string{}
	for _, i := range array1 {
		for _, j := range array2 {
			if !stringinslice.StringInSlice(i, newString) && strings.Contains(j, i) {
				newString = append(newString, i)
			}
		}
	}
	sort.Strings(newString)
	return newString
}
