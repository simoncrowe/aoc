package parseutil

import (
	"log"
	"strconv"
	"strings"
)

func ParseInt(val string) int {
	trimmed := strings.Trim(val, " ")
	parsed, err := strconv.Atoi(trimmed)
	if err != nil {
		log.Fatal(err)
	}
	return parsed
}

func ParseInts(vals []string) []int {
	ints := []int{}
	for _, val := range vals {
		parsed := ParseInt(val)
		ints = append(ints, parsed)
	}
	return ints
}
