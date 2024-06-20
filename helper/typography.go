package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func StringSplitPure(data string, sep string) []string {
	var result []string
	
	split := strings.Split(data, sep)
	
	for _, v := range split {
		if len(strings.TrimSpace(v)) > 0 {
			result = append(result, v)
		}
	}
	
	return result
}

func UniqueString(elements []string) []string {
	encountered := map[string]bool{}
	
	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}
	
	// Place all keys from the map into a slice.
	var result []string
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func GenerateRandomAlphanumeric(n int) string {
	if n == 0 {
		n = 6
	}
	
	const (
		letterBytes   = "abcdefghijklmnopqrstuvwxyz1234567890"
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	src := rand.NewSource(time.Now().UnixNano())
	
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	
	return string(b)
}

func UnderscoreToCamelCase(s string) string {
	return strings.Replace(strings.Title(strings.Replace(strings.ToLower(s), "_", " ", -1)), " ", "", -1)
}

func CamelCaseToUnderscore(str string) string {
	var output []rune
	var segment []rune
	
	for _, r := range str {
		// not treat number as separate segment
		if !unicode.IsLower(r) && string(r) != "_" && !unicode.IsNumber(r) {
			output = addSegment(output, segment)
			segment = nil
		}
		segment = append(segment, unicode.ToLower(r))
	}
	output = addSegment(output, segment)
	
	return string(output)
}

func CensorName(inputStr string) string {
	substrings := strings.Split(inputStr, " ")
	
	for i := range substrings {
		substrings[i] = CensorSubstring(substrings[i])
	}
	
	outputStr := strings.Join(substrings, " ")
	return outputStr
}

func CensorSubstring(substring string) string {
	if len(substring) == 1 {
		return substring
	}
	return string(substring[0]) + strings.Map(
		func(r rune) rune {
			if r == rune(substring[0]) || r == rune(substring[len(substring)-1]) {
				return r
			}
			return '*'
		}, substring[1:len(substring)-1],
	) + string(substring[len(substring)-1])
}

func CensorNominal(num int) string {
	numStr := fmt.Sprintf("%d", num)
	n := len(numStr)
	
	// Replace each digit with the character 'x'
	for i := 0; i < n; i++ {
		numStr = numStr[:i] + "X" + numStr[i+1:]
	}
	
	// Insert a dot separator after every third character
	if n > 3 {
		for i := n - 3; i > 0; i -= 3 {
			numStr = numStr[:i] + "." + numStr[i:]
		}
	}
	
	return numStr
}
