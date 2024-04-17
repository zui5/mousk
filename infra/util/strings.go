package util

import "sort"

func Uint32SliceEqual(src []uint32, dst []uint32) bool {
	if len(src) != len(dst) {
		return false
	}

	sort.Slice(src, func(i, j int) bool {
		return src[i] >= src[j]
	})

	sort.Slice(dst, func(i, j int) bool {
		return dst[i] >= dst[j]
	})

}
