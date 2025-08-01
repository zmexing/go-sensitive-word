package go_sensitive_word

import _ "embed"

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

var (
	//go:embed text/COVID-19词库.txt
	DictCovid19 string
	//go:embed text/其他词库.txt
	DictOther string
	//go:embed text/反动词库.txt
	DictReactionary string
	//go:embed text/暴恐词库.txt
	DictViolence string
	//go:embed text/民生词库.txt
	DictPeopleLife string
	//go:embed text/色情词库.txt
	DictPornography string
	//go:embed text/补充词库.txt
	DictAdditional string
	//go:embed text/贪腐词库.txt
	DictCorruption string
	//go:embed text/零时-Tencent.txt
	DictTemporaryTencent string
)
