package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SayHello() {
	fmt.Println("はろー")
}

func GetRequestIndex(query string) int {
	re, _ := regexp.Compile("index=([0-9]+)")
	result := re.FindSubmatch([]byte(query))

	if len(result) > 1 {
		q := string(result[1])
		if num, err := strconv.Atoi(q); err == nil {
			return num
		}
	}

	return 0
}

func IsExtinf(line string) bool {
	if m, _ := regexp.MatchString("^#EXTINF:[0-9]+(\\.[0-9]+)?,$", line); !m {
		return false
	}
	return true
}

func IsTsFile(line string) bool {
	return strings.Contains(line, ".ts")
}
