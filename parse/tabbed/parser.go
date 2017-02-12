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

/*
 Go's time package provides its users with a distinctive mechanism for
 specifying how datetime strings should be formatted.

 If you are used to strftime/strptime usage with formatters that look like:
 "%Y%m%d %H:%M:%S", then prepare for a shock.

 In Go you take the magic numbers 1-7 and use them to write out a datetime
 according to the string layout of your choice.

 Mnemonic for the layout: 1, 2, 3, 4, 5, 6, 7
 1(month) 2(day) 3(pm) 4(min) 5(sec) 6(year) 7(timezone), as in:

      "01/02 03:04:05PM '06 -0700"

 We want RFC 3339 without any fractional seconds, hence:
      "2006-01-02T15:04:05Z"

  http://stackoverflow.com/a/25845833/10278
*/
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
	if !strings.HasPrefix(versionString, versionPrefix) {
		panicMsg := "Invalid file. "
		panicMsg += "First content line must contain "
		panicMsg += "version code starting with prefix: %s"
		panic(fmt.Sprintf(panicMsg, versionPrefix))
	}
}

func makeItemFromLine(line string) *item.Item {
	pieces := strings.Split(line, "\t")
	if len(pieces) != columnsPerLine {
		panicMsg := "Invalid line layout. Needs %d columns. Line: %s"
		panic(fmt.Sprintf(panicMsg, columnsPerLine, line))
	}
	return item.NewWithGivenTime(toId(pieces[0]),
		toActionType(pieces[1]),
		toDesc(pieces[2]),
		toTime(pieces[3]))
}

func toId(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		panicMsg := "Invalid ref id. Must be an integer. Line: %s"
		panic(fmt.Sprintf(panicMsg, text))
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
		panicMsg := "Invalid datetime. Time parse error. %v. Line: %s"
		panic(fmt.Sprintf(panicMsg, err, text))
	}

	return t
}
