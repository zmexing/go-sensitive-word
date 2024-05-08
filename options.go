package go_sensitive_word

const (
	StoreMemory = iota
)

const (
	FilterDfa = iota
)

type StoreOption struct {
	Type uint32
}

type FilterOption struct {
	Type uint32
}
