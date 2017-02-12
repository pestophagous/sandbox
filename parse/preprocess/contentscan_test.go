package preprocess

import (
	"testing"

	"fmt"
	"io"

	localtests "github.com/practicum/sandbox/testing"
)

var sampleTxtFilepath string = localtests.
	DataAssetFullPath("january.txt", reflectionToy{})

// reflectionToy is a dummy type created in the preprocess package as a trick to
// retrieve the package name. http://stackoverflow.com/a/25263604/10278
type reflectionToy struct{}

func TestScan(t *testing.T) {
	datapath := localtests.DataAssetFullPath("january.txt", reflectionToy{})
	iterator := NewContentIterator(datapath)

	_, err := iterator()

	for err == nil {
		_, err = iterator()
	}

	if err != io.EOF {
		t.Error(err)
	}
}

func ExampleNewContentIterator() {
	iterator := NewContentIterator(sampleTxtFilepath)

	line, err := iterator()

	for err == nil {
		fmt.Println(line)
		line, err = iterator()
	}

	if err != io.EOF {
		panic(fmt.Sprintf("Error while iterating over %s.", sampleTxtFilepath))
	}
}
