package item

import (
	"fmt"
	"time"
)

type ActionType int

const (
	Open ActionType = iota
	Close
)

// ActionType.String method fulfills 'Stringer interface' from Go 'fmt' package.
func (t ActionType) String() string {
	if t == Open {
		return "++"
	} else if t == Close {
		return "--"
	} else {
		return "??error??"
	}
}

type Item struct {
	id      int
	action  ActionType
	desc    string
	timeUTC time.Time
}

func NewWithGivenTime(refid int,
	actionType ActionType,
	description string,
	timeInUTC time.Time) *Item {

	return &Item{
		id:      refid,
		action:  actionType,
		desc:    description,
		timeUTC: timeInUTC,
	}
}

func NewAtTimeNow(refid int, actionType ActionType, description string) *Item {
	return &Item{
		id:      refid,
		action:  actionType,
		desc:    description,
		timeUTC: time.Now().UTC(),
	}
}

// Item.String method fulfills 'type Stringer interface' from Go 'fmt' package.
func (i Item) String() string {
	return fmt.Sprintf("%s %s %d: %s",
		i.LocalDatetimeStr(),
		i.action,
		i.id,
		i.desc)
}

func (i Item) LocalDatetimeStr() string {
	return i.timeUTC.Local().Format(time.Stamp)
}

func (i Item) ID() int {
	return i.id
}

func (i Item) ActionType() ActionType {
	return i.action
}
