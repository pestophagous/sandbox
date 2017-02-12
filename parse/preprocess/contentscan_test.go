package preprocess

import (
	"testing"

	"fmt"
	"os"
	"path"
)

func TestScan(t *testing.T) {
	datapath := os.Getenv("GOPATH")
	datapath = path.Join(datapath, "src", "github.com", "practicum", "sandbox", "sampledata", "january.txt")
	iterator := scan(datapath)

	line, err := iterator()

	for err == nil {
		line, err = iterator()
		fmt.Println(line)
	}

	t.Error(err)
}
