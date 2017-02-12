package tabbed

import (
	"testing"

	"fmt"

	localtests "github.com/practicum/sandbox/testing"
)

// a dummy type to help retrieve the package name
//    with thanks to: http://stackoverflow.com/a/25263604/10278
type reflectionToy struct{}

func TestParse(t *testing.T) {
	const targetFile = "january.txt"
	datapath := localtests.DataAssetFullPath(targetFile, reflectionToy{})
	results, _ := Parse(datapath)

	if len(results) == 0 {
		t.Errorf("empty results while parsing %s", targetFile)
	}

	for i, r := range results {
		fmt.Printf("line:%d\n", i+1)
		fmt.Println(r)
	}
}
