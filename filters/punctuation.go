package filters

import (
	"regexp"
	"strings"
)

type puctuation struct {
}

// NewPunctuation фильтр знаков пунктуации
func NewPunctuation() TextFilter {
	return puctuation{}
}

func (p puctuation) Filter(s string) string {
	r := strings.NewReplacer(
		" ", "",
		", ", " ",
		". ", " ",
		"! ", " ",
		"? ", " ",
		"–", " ",
	)

	ret := r.Replace(s)

	re := regexp.MustCompile(`(\!|\.|\?)`)
	ret = re.ReplaceAllString(ret, "")

	return ret
}
