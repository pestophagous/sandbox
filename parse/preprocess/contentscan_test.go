package preprocess

import (
	"testing"

	"fmt"
	"io"

	localtests "github.com/practicum/sandbox/testing"
)

// reflectionToy is a dummy type created in the preprocess package as a trick to
// retrieve the package name. http://stackoverflow.com/a/25263604/10278
type reflectionToy struct{}

func TestScan(t *testing.T) {
	datapath := localtests.DataAssetFullPath("january.txt", reflectionToy{})
	iterator := NewContentIterator(datapath)

	line, err := iterator()

	for err == nil {
		line, err = iterator()
		fmt.Println(line)
	}

	if err != io.EOF {
		t.Error(err)
	}
}
