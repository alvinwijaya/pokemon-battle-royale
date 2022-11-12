package helper

import "strconv"

func ConvertStringToUint64(data string) uint64 {
	res, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0
	}
	return res
}
