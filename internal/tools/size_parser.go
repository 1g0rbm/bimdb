package tools

import (
	"fmt"
	"strconv"
)

func ParseSize(str string) (int, error) {
	if len(str) == 0 || str[0] < '0' || str[0] > '9' {
		return 0, fmt.Errorf("invalid input size format")
	}

	idx := 0
	size := 0

	for idx < len(str) && str[idx] >= '0' && str[idx] <= '9' {
		num, err := strconv.Atoi(string(str[idx]))
		if err != nil {
			return 0, err
		}

		size = size*10 + num
		idx += 1
	}

	units := str[idx:]
	switch units {
	case "B", "b", "":
		return size, nil
	case "KB", "Kb", "kb":
		return size * 1000, nil
	case "MB", "Mb", "mb":
		return size * 1000 * 1000, nil
	case "GB", "Gb", "gb":
		return size * 1000 * 1000 * 1000, nil
	default:
		return 0, fmt.Errorf("incorrect size")
	}
}
