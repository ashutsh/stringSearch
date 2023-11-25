package utils

import "strconv"


// Takes in a string slice containing all integers in string and converts them to int and returns a []int. 
// Also as of now, it does not handle error so if the string slice contains any deformity, the function might not work as expected without handling error.
func ConvertStringToIntSlice(strslc []string) []int {
	result := make([]int, 0)
	for _, r := range strslc {
		rint, _ := strconv.Atoi(r) 
		result = append(result, rint)
	}

	return result
}


// Takes in an Int slice and returns an Int slice containing no repeat characters. 
// Also, the first occurence of the integer/element will be kept and rest will be removed. 
func UniqueIntSlc(dataslc []int) []int {
	resultSlc := make([]int, 0)
	datamap := make(map[int]bool)
	for _, r := range dataslc {
		if k := datamap[r] ; !k {
			datamap[r] = true
			resultSlc = append(resultSlc, r)
		} 
	}

	return resultSlc
}