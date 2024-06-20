package utils

import (
	"reflect"
	"strconv"
	"strings"
)

// Append ...
func Append(slice []interface{}, items ...interface{}) []interface{} {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

// Extend ...
func Extend(slice []interface{}, element interface{}) []interface{} {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]interface{}, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// ConvertSliceStringToInterface ...
func ConvertSliceStringToInterface(data []string) []interface{} {
	newData := make([]interface{}, len(data))
	
	for i, v := range data {
		newData[i] = v
	}
	
	return newData
}

// Set setup args for query param
func Set(val reflect.Value) map[string]interface{} {
	queries := make(map[string]interface{})
	
	for i := 0; i < val.NumField(); i++ {
		isEmpty := true
		q := val.Type().Field(i).Tag.Get("query")
		value := reflect.Indirect(val).FieldByName(val.Type().Field(i).Name).Interface()
		
		if q != "" {
			switch value.(type) {
			case int:
				if value.(int) > 0 {
					isEmpty = false
					value = strconv.Itoa(value.(int))
				}
				
				break
			case string:
				if len(value.(string)) > 0 {
					isEmpty = false
					value = strings.Replace(value.(string), "'", "''", -1)
				}
				
				break
			default:
				break
			}
			
			if !isEmpty {
				queries[q] = value
			}
		}
	}
	
	return queries
}
