package utils

import (
	"strconv"
)

func Abs[T int | int8 | int16 | int32 | int64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Atoi[T string | rune](x T) int {
	xStr := string(x)
	num, _ := strconv.Atoi(xStr)
	return num
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return Abs(a) * (Abs(b) / GCD(a, b))
}

func Min[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | float32 | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Pow2[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](exp T) T {
	return (1 << exp)
}
