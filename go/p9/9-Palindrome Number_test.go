package p9

import "testing"

// -121 false -> 所有负数都false

// 0 true -> 奇数个
// 121 true

// 11 true -> 偶数

// 思路1：字符串，两边向中间对比
// 思路2：每个数搞成数组，两边向中间对比
// 思路3：反转整数，判断是否相等，反转可能有溢出，所以提高到int64，思路3比1、2节省空间
// 思路4：输入是对称的，翻转后一定也不会溢出，反转如果溢出，直接返回false，思路4更节省空间、时间
// 思路5：反转一半，如果刚好一半的时候，2个相等，则true，如果反转的那部分大于剩下部分，
// 说明超过一半了还没有相等，返回false。反转的部分和剩下的部分对比时，也要注意偶数和奇数长度2种情况。

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	// x > 0
	return method5(x)
}

func method5(x int) bool {
	var revert int
	for x > revert {
		revert = revert*10 + x%10
		x = x / 10
	}
	return x == revert || x == revert/10
}

func method4(x int) bool {
	oldX := x
	var max int = (1 << 31) - 1
	var min int = -(1 << 31)
	var ret int
	for x != 0 {
		num := x % 10
		if ret > max/10 {
			return false
		}
		if ret < min/10 {
			return false
		}

		ret = ret*10 + num
		x = x / 10
	}
	return oldX == ret
}

func method3(x int) bool {
	var x64 = int64(x)
	ret := reverse64(x64)
	return x64 == ret
}

func reverse64(x int64) int64 {
	var ret int64
	for x != 0 {
		num := x % 10
		ret = ret*10 + num
		x = x / 10
	}
	return ret
}

func TestPalin(t *testing.T) {
	cs := []struct {
		in  int
		exp bool
	}{
		{0, true},
		{1, true},
		{9, true},
		{11, true},
		{121, true},
		{-1, false},
		{-129283, false},
		{123, false},
		{12, false},
	}

	for i, c := range cs {
		got := isPalindrome(c.in)
		if got != c.exp {
			t.Errorf("%d, %v, want %v, got %v", i, c.in, c.exp, got)
		}
	}
}
