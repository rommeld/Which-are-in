package whicharein

import (
	"sort"
	"stringinslice"
	"strings"
)

var newString, sortedString []string

//InArray
func InArray(array1 []string, array2 []string) []string {
	for _, i := range array2 {
		for _, j := range array1 {
			if !stringinslice.StringInSlice(i, newString) && strings.Contains(i, j) {
				newString = append(newString, j)
			}
		}
	}
	sortedString = sort.Sort(newString)
	return sortedString
}
