package string_camelcase

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	reg := regexp.MustCompile(`[_|-].`)

	result := reg.ReplaceAllStringFunc(s, func(match string) string {
		return strings.ToUpper(string(match[1]))
	})

	fmt.Println("string: ", s, "result: ", result)

	return result
}
