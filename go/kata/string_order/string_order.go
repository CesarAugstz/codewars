package string_order

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Order(sentence string) (res string) {
	if sentence == "" {
		return sentence
	}

	parts := strings.Split(sentence, " ")

	ordered := make([]string, len(parts))

	for _, v := range parts {
		pos, err := strconv.Atoi((regexp.MustCompile(`\d`).FindString(v)))
		if err != nil {
			panic("Cannot convert the index to integer" + fmt.Sprint(err))
		}

		ordered[pos-1] = v
	}

	return strings.Join(ordered, " ")
}
