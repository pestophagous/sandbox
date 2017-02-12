package tabbed

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/practicum/sandbox/core/item"
	"github.com/practicum/sandbox/parse/preprocess"
)

const columnsPerLine = 4

// Mnemonic for the layout cue: 1, 2, 3, 4, 5, 6, 7
// 1(month) 2(day) 3(pm) 4(min) 5(sec) 6(year) 7(timezone)
// 01/02 03:04:05PM '06 -0700
// http://stackoverflow.com/a/25845833/10278
const requiredTimeLayout = "2006-01-02T15:04:05Z"

func Parse(pathtofile string) ([]*item.Item, error) {
	var firstline string
	var results []*item.Item

	iterator := preprocess.NewContentIterator(pathtofile)

	line, err := iterator()

	if err == nil {
		firstline = line
		line, err = iterator()
	}

	// only check version once we know we HAVE a second line
	if err == nil {
		panicIfVersionLineIsBad(firstline)

		results = append(results, makeItemFromLine(line))
		line, err = iterator()
	}

	for err == nil {
		results = append(results, makeItemFromLine(line))
		line, err = iterator()
	}

	if err == io.EOF {
		err = nil
	}
	return results, err
}

func panicIfVersionLineIsBad(versionString string) {
	const versionPrefix = "v"
	panicMsg := "Invalid file. "
	panicMsg += "First content line must contain version code starting with prefix: %s"
	if !strings.HasPrefix(versionString, versionPrefix) {
		panic(fmt.Sprintf(panicMsg, versionPrefix))
	}
}

func makeItemFromLine(line string) *item.Item {
	pieces := strings.Split(line, "\t")
	if len(pieces) != columnsPerLine {
		panic(fmt.Sprintf("Invalid line layout. Needs %d columns. Line: %s", columnsPerLine, line))
	}
	return item.NewWithGivenTime(toId(pieces[0]),
		toActionType(pieces[1]),
		toDesc(pieces[2]),
		toTime(pieces[3]))
}

func toId(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		panic(fmt.Sprintf("Invalid ref id. Must be an integer. Line: %s", text))
	}

	return i
}

func toActionType(text string) item.ActionType {
	if text == "open" {
		return item.Open
	} else if text == "close" {
		return item.Close
	} else {
		panic(fmt.Sprintf("Invalid action type. Line: %s", text))
	}
}

func toDesc(text string) string {
	// no transformation needed at the present time
	return text
}

func toTime(text string) time.Time {
	t, err := time.Parse(requiredTimeLayout, text)
	if err != nil {
		panic(fmt.Sprintf("Invalid datetime. Time parse error. %v. Line: %s", err, text))
	}

	return t
}
