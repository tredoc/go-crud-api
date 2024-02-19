package repository

import (
	"strconv"
	"strings"
)

func stringToInt64Slice(s string) ([]int64, error) {
	s = strings.Trim(s, "{}")
	if s == "" {
		return []int64{}, nil
	}
	parts := strings.Split(s, ",")
	result := make([]int64, len(parts))
	for i, part := range parts {
		n, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = n
	}
	return result, nil
}
