package filters

import (
	"strings"
)

type brackets struct {

}

func NewBrackets() TextFilter {
	return brackets{}
}

func (p brackets) Filter(s string) string {
	r := strings.NewReplacer(
		"(", "",
		")", "",
		"[", "",
		"]", "",
		"{", "",
		"}", "",
	)

	ret := r.Replace(s)

	return ret
}
