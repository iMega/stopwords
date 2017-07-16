package filters

import "regexp"

type doublecates struct {
}

func NewDoublecates() TextFilter {
	return &doublecates{}
}

func (f doublecates) Filter(s string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}
