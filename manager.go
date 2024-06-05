package go_sensitive_word

import (
	"errors"
	"github.com/zmexing/go-sensitive-word/filter"
	"github.com/zmexing/go-sensitive-word/store"
)

type Manager struct {
	store.Store
	filter.Filter
}

func NewFilter(storeOption StoreOption, filterOption FilterOption) (*Manager, error) {
	var filterStore store.Store
	var myFilter filter.Filter

	switch storeOption.Type {
	case StoreMemory:
		filterStore = store.NewMemoryModel()
	default:
		return nil, errors.New("invalid store type")
	}

	switch filterOption.Type {
	case FilterDfa:
		dfaModel := filter.NewDfaModel()
		go dfaModel.Listen(filterStore.GetAddChan(), filterStore.GetDelChan())
		myFilter = dfaModel
	default:
		return nil, errors.New("invalid filter type")
	}

	return &Manager{
		Store:  filterStore,
		Filter: myFilter,
	}, nil
}
