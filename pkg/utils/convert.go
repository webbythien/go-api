package utils

import (
	"log"
	"strconv"
)

func String2Int64(num string) int64 {
	number, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatal("Couldn't parse number from string")
	}
	return number
}
