package json_to_go

import "strings"

func Intersect(a []string, b []string) (res []string) {
	book := make(map[string]struct{})
	for _, i := range a {
		book[i] = struct{}{}
	}
	for _, i := range b {
		if _, ok := book[i]; ok {
			res = append(res, i)
		}
	}
	return
}

func ToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}
