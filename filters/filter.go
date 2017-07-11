package filters

type TextFilter interface {
	Filter(string string) string
}
