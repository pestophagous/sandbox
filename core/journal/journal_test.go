package journal

import (
	"testing"

	"fmt"

	"github.com/practicum/sandbox/core/item"
	localtests "github.com/practicum/sandbox/testing"
)

var janTxtFile *Journal

func init() {
	datafile := localtests.DataAssetFullPath("january.txt", Journal{})
	var err error
	janTxtFile, err = OpenFromFile(datafile)
	if err != nil {
		panic(fmt.Sprintf("Can't load %s", datafile))
	}
}

func TestScan(t *testing.T) {
	fmt.Printf("Completed: %d\n", janTxtFile.CountCompletedItems())
}

func TestSlicing(t *testing.T) {
	some := janTxtFile.GetCompletedItems()
	for _, item := range some {
		fmt.Println(item)
	}

	// if GetCompletedItems returned an aliased slice that shared an underlying
	// array with the Journal, this could wreak havoc:
	some[0] = &item.Item{}

	// Were it not for the immutability of Item, we would also risk:
	// some[3].CorruptIt()

	smore := janTxtFile.GetCompletedItems()
	for _, item := range smore {
		fmt.Println(item)
	}
}
