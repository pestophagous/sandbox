package preprocess

import (
	"testing"

	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

func TestScan(t *testing.T) {
	datapath := os.Getenv("GOPATH")
	datapath = path.Join(datapath, "src", "github.com", "practicum", "sandbox", "sampledata", "january.txt")
	iterator := scan(datapath)

	io.EOF = errors.New("vaaarf")

	line, err := iterator()

	for err == nil {
		line, err = iterator()
		fmt.Println(line)
	}

	t.Error(err)
}
