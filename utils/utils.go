package utils

import (
	"crypto/md5"
	"fmt"
)

func CalPage(count int64, page, size int) (int, int) {
	if size <= 0 {
		size = 10
	}
	if page < 1 {
		page = 1
	}

	total := (int(count) + size - 1) / size
	if page > total {
		page = total
	}

	offset := (page - 1) * size

	return size, offset
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
