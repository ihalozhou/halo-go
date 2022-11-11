package _array

import (
	"fmt"
	"strconv"
)

func InArray(value interface{}, array []string) bool {
	for _, v := range array {
		fmt.Printf("%#v v: %#v value:\n ", v, value)
		if v == value {
			return true
		}
	}
	return false
}

func InIntArray(value int, array []int) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func StringToInt(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}
