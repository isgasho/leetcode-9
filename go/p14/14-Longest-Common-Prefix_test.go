package p14

import "testing"

// 思路1：先计算出最短的字符串，然后同时遍历每一个字符串，如果相同下标相同，
// 则为共同前缀，否则不是，退出返回最长公共前缀。
// 思路2：从前到后，每次都找2个字符串的lcp，最终得到所有字符串的lcp。
// 思路3：分治法，把字符串数组分成2半字符串数组，每个字符串递归去查找lcp，
// 最后再两两归并得到最后的lcp。
// 思路4：二分查找法，新找到最小的长度，拿最短的字符串进行二分查找，判断前一半是不是所有字符串的cp，
// 是则向后找，不是则向前找，直到找到边界，得到最终lcp的长度L，则前L个字符就是lcp。

// 思路1的空间和时间最优，时间O(S)，空间O(1)。
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	return method1(strs)
}

func method1(strs []string) string {
	l := len(strs[0])
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < l {
			l = len(strs[i])
		}
	}

	prefix := ""
	for i := 0; i < l; i++ {
		c := strs[0][i]
		same := true
		for j := 0; j < len(strs); j++ {
			if c != strs[j][i] {
				same = false
				break
			}
		}

		if same {
			prefix = prefix + string(c)
		} else {
			return prefix
		}
	}
	return prefix
}
func TestLCP(t *testing.T) {
	cs := []struct {
		in  []string
		exp string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{"flower", "flow", "fl"},
			"fl",
		},
		{
			[]string{"dog", "race", "dag"},
			"",
		},
		{
			[]string{"dog", ""},
			"",
		},
	}

	for i, c := range cs {
		got := longestCommonPrefix(c.in)
		if got != c.exp {
			t.Errorf("case %d, in: %v, want: %v, got: %v", i, c.in, c.exp, got)
		}
	}
}
