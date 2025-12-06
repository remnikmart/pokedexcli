package funcs

import "strings"

func CleanInput(text string) []string {
	res := []string{}
	text = strings.Trim(strings.Trim(text, " "), "\n")
	if text == "" {
		return res
	}
	for v := range strings.SplitSeq(text, " ") {
		s := strings.Trim(v, " ")
		if s != "" {
			res = append(res, strings.ToLower(s))
		}
	}
	return res
}
