package lengthoflongestsubstring

import (
	"fmt"
	"strings"
	"testing"
)

func TestXxx(t *testing.T) {
	count := lengthOfLongestSubstring("123412345611123")
	fmt.Println(count)
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		fmt.Println(s[start:i])
		fmt.Println(index, string(s[i]),i)
		if index == -1 && (i+1) > end {
			end = i + 1
		} else {
			start += index + 1
			end += index + 1
			fmt.Println("start,end",start,end)
		}
	}
	return end - start
}
