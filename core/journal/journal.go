package journal

import (
	"github.com/practicum/sandbox/core/item"
	_ "github.com/practicum/sandbox/parse/json"
	"github.com/practicum/sandbox/parse/tabbed"
)

type Journal struct {
	items   []*item.Item
	buckets bucketsCache
}

func OpenFromFile(pathAndName string) (*Journal, error) {
	allitems, err := tabbed.Parse(pathAndName)
	if err != nil {
		return nil, err
	}

	j := &Journal{
		items: allitems,
	}
	j.buckets = newCache(j)

	return j, nil
}

func CreateInDirectory(dirpath string) *Journal { // TODO implement
	return &Journal{}
}

func (j *Journal) CountCompletedItems() int {
	j.buckets.update()
	return len(j.buckets.completed)
}

func (j *Journal) GetCompletedItems() []*item.Item {
	j.buckets.update()

	// Do _NOT_ return 'j.buckets.completed' directly.  That would allow the
	// caller to OVERWRITE Item pointers in the slice, affecting/corrupting the
	// j.buckets cache!

	// Make an entirely new, safe slice to return.
	// (it will point to the same Item pointers, but Item(s) are immutable)
	result := make([]*item.Item, len(j.buckets.completed))
	copy(result, j.buckets.completed)
	return result
}
