package preprocess

import (
	"bufio"
	"io"
	"os"
)

// ... returns a closure (a function). Call the function repeatedly until an error is returned.
// If you do not exhaust calls until io.EOF or some other error, the file will not be closed.
func scan(filename string) func() (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return func() (string, error) {
			return "", err
		}
	}

	scanner := bufio.NewScanner(file)

	return func() (string, error) {
		if err != nil {
			return "", err
		}

		if scanner.Scan() {
			// still getting lines. return the line:
			return scanner.Text(), nil
		} else {
			// nothing left to scan. arrage to close file, and set fixed err value.
			defer file.Close()

			// did we end normally?
			if err = scanner.Err(); err == nil {
				// normal case. replace scanner default nil with EOF:
				err = io.EOF
				return "", err
			} else {
				// scanner gave a non-nil error, so use it:
				return "", err
			}
		}
	}
}
