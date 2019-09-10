package p7

import "testing"

// 思路1：使用int64做转换和判断溢出，溢出时，转为int为0
// 思路2：使用int做转换，但提前判断溢出

// 注意：在Go中，负数的取余自动带负数，所以计算完结果自动为负数，所以
// 不需要区分正数还是负数

// 已经保证入参，所以无需校验
// 提前判断是否会有溢出
// 因为int最大值和最小值的第一位都是2，翻转后必定在[-8,7]区间，
// 所以简化溢出判断，删除（ret == max/10 && ret > 7)，
// (ret == min/10 && ret < -8)

func reverse(x int) int {
	var max int = (1 << 31) - 1
	var min int = -(1 << 31)
	var ret int
	for x != 0 {
		num := x % 10
		if ret > max/10 {
			return 0
		}
		if ret < min/10 {
			return 0
		}

		ret = ret*10 + num
		x = x / 10
		print(ret, "\t", x, "\n")
	}

	return ret
}

func TestReverse(t *testing.T) {
	type cases struct {
		in  int
		exp int
	}

	cs := []cases{
		{1, 1},
		{123, 321},
		{0, 0},
		{199999999, 999999991},
		{2147483647, 0},
		{1563847412, 0},
		{-1, -1},
		{-123, -321},
		{-199999999, -999999991},
		{-2147483648, 0},
	}

	for i, c := range cs {
		got := reverse(c.in)
		if got != c.exp {
			t.Errorf("case %d: input: %d, want: %d, got: %d", i, c.in, c.exp, got)
		}
	}
}
