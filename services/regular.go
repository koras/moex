package services

import (
	"regexp"
	"strings"
)

func ClearText(content string) string {
	//	re := regexp.MustCompile("JSON_CALLBACK\\(")
	//replaced := re.ReplaceAllString(content, "")

	re3 := regexp.MustCompile("\\{\"charsetinfo\"\\: \\{\"name\": \"utf-8\"\\}\\},")
	replaced3 := re3.ReplaceAllString(content, "")
	re2 := regexp.MustCompile("\\)$")
	replaced2 := re2.ReplaceAllString(replaced3, "")
	replaced4 := strings.TrimSpace(replaced2)
	re4 := regexp.MustCompile("^\\[")
	replaced5 := re4.ReplaceAllString(replaced4, "")
	re5 := regexp.MustCompile("\\]$")
	contentReplace := re5.ReplaceAllString(replaced5, "")

	//	fmt.Print(contentReplace)

	return contentReplace
}
