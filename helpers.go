package grug

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// anything to string
func atostr(a interface{}) string {
	return fmt.Sprintf("%v", a)
}

// Applies a function to every element of a value assumed to be a slice
func fMapSlice(val interface{}, mapFunc func(interface{})) {
	reflectionVal := reflect.ValueOf(val)
	if reflectionVal.Kind() == reflect.Slice || reflectionVal.Kind() == reflect.Array {
		for i := 0; i < reflectionVal.Len(); i++ {
			mapFunc(reflectionVal.Index(i).Interface())
		}
	}
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

// fieldAccess accesses fields of an object according to a plan
// Supports both indexing and named access. E.g. FieldA.0.B would access args[0].FieldA[0].B
func fieldAccess(object interface{}, accessPattern string) (interface{}, error) {
	if accessPattern == "." {
		return object, nil
	}
	indexPlan := strings.Split(atostr(accessPattern), ".")

	var value interface{}
	for n, field := range indexPlan {
		if n == 0 {
			value = object
		}

		idx, err := strconv.Atoi(field)
		if err == nil {
			valAsSlice, ok := value.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%v is not accessible by index", value)
			}
			if idx > len(valAsSlice)-1 {
				return nil, fmt.Errorf("%d is out of range for slice %v with length %d", idx, value, len(valAsSlice))
			}
			value = valAsSlice[idx]
		} else {
			valAsMap, ok := value.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v is not accessible by name %s", value, field)
			}
			value = valAsMap[field]
		}
	}

	return value, nil
}
