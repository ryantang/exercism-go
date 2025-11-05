package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`^(\[FTL\]|\[ERR\]|\[WRN\]|\[INF\]|\[TRC\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`".*(?i)password.*"`)

	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`User\s+(\w+)`)
	for index := range lines {
		username := re.FindStringSubmatch(lines[index])
		if username != nil {
			lines[index] = fmt.Sprintf("[USR] %s %s", username[1], lines[index])
		}
	}
	return lines
}
