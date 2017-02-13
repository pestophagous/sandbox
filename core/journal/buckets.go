package journal

import (
	"fmt"

	itemp "github.com/practicum/sandbox/core/item"
)

const itemTypeOpenIsAlwaysFirst = "For any refID, the first journal item seen must be of type item.Open"

// in its current nascent state, we basically recompute the whole cache
// whenever we need a value from it, so it isn't yet much of a cache ;)
type bucketsCache struct {
	matchmap   mapItemToAction
	journal    *Journal
	incomplete []*itemp.Item
	completed  []*itemp.Item
}

// maps an Item ID to the 'winning action' (last action wins)
type mapItemToAction map[int]*itemWithWinningAction

type itemWithWinningAction struct {
	item          *itemp.Item
	winningAction itemp.ActionType
}

func newCache(j *Journal) bucketsCache {
	return bucketsCache{
		matchmap: make(mapItemToAction),
		journal:  j,
	}
}

// as stated above, we currently 'update the world' (total cache invalidation
// every time we touch the cache).  this can definitely be improved.
func (b *bucketsCache) update() {
	// start by resetting everything
	b.matchmap = make(mapItemToAction)
	b.incomplete = []*itemp.Item{}
	b.completed = []*itemp.Item{}

	for _, item := range b.journal.items {
		if val, ok := b.matchmap[item.ID()]; !ok {
			// no prior map info for this item.Item
			b.matchmap[item.ID()] = &itemWithWinningAction{
				item:          item,
				winningAction: item.ActionType(),
			}
			panicIfUnexpectedAT(item, itemp.Open, itemTypeOpenIsAlwaysFirst)
		} else {
			// we already saw the item, now we must only update the action:
			val.winningAction = item.ActionType()
		}
	}

	// TODO deal with random order of map iteration
	for _, v := range b.matchmap {
		if v.winningAction == itemp.Close {
			b.completed = append(b.completed, v.item)
		} else {
			b.incomplete = append(b.incomplete, v.item)
		}
	}
}

func panicIfUnexpectedAT(item *itemp.Item, action itemp.ActionType, msg string) {
	if item.ActionType() != action {
		fmt.Println(item)
		panic(msg)
	}
}
