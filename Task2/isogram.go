package isogram

import "strings"

func IsIsogram(str string) bool {
	lowStr := strings.ToLower(str)
	for i := 0; i < len(lowStr); i++ {
		if lowStr[i] != ' ' && lowStr[i] != '-' {
			for j := i + 1; j < len(lowStr); j++ {
				if lowStr[i] == lowStr[j] {
					return false
				}
			}
		}
	}
	return true
}
