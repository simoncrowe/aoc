package parseutil

import (
	"log"
	"strconv"
)

func ParseInt(val string) int {
	parsed, err := strconv.Atoi(val)
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
