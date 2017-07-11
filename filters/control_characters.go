package filters

import "regexp"

type control_characters struct {

}

func NewControlCharacters() TextFilter {
	return &control_characters{}
}

func (f control_characters) Filter(s string) string {
	re := regexp.MustCompile(`(\r|\n|\t|\r)`)
	ret := re.ReplaceAllString(s, " ")

	return ret
}
