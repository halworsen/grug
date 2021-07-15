package grug

import (
	"fmt"
	"reflect"
	"strconv"
)

// anything to string
func atostr(a interface{}) string {
	return fmt.Sprintf("%v", a)
}

// appends a value but "opens up" slices and arrays and adds every element instead of the slice itself
func appendExpandSlices(slice []interface{}, val interface{}) []interface{} {
	reflectionVal := reflect.ValueOf(val)
	if reflectionVal.Kind() == reflect.Slice || reflectionVal.Kind() == reflect.Array {
		for i := 0; i < reflectionVal.Len(); i++ {
			slice = append(slice, reflectionVal.Index(i).Interface())
		}
	} else {
		slice = append(slice, val)
	}
	return slice
}

// Given an array of the form ["lower" ":" "upper"], returns the lower/upper bounds as ints
func getSliceBounds(slice []string, max int) (int, int) {
	lower, upper := 0, max

	if len(slice) < 3 {
		return lower, upper
	}

	if slice[0] != "" {
		lowerIn, err := strconv.Atoi(slice[0])
		if lowerIn < 0 {
			lowerIn += max
		} else {
			lowerIn -= 1
		}
		if err == nil {
			lower = lowerIn
		}
	}

	if slice[1] != "" {
		upperIn, err := strconv.Atoi(slice[2])
		if upperIn < 0 {
			upperIn += max
		} else {
			upperIn -= 1
		}
		if err == nil {
			upper = upperIn
		}
	}

	return lower, upper
}
