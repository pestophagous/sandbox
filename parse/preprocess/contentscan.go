package preprocess

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const commentPrefix = "//"

// NewContentIterator returns a closure (a function). Call the function
// repeatedly to retrieve preprocessed lines until an error is returned.  The
// underlying os.File won't be closed until you exhaust the iterator and reach
// the io.EOF condition.
func NewContentIterator(filename string) func() (string, error) {

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

		if text, err := getNextContentRichLine(scanner); err == nil {
			// still getting lines. return the line:
			return text, nil
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

// isContentRich checks a line for blankness and comment-ness.
// Anything without those properties is deemed "content rich".
func isContentRich(line string) bool {
	line = strings.TrimSpace(line)
	isNotBlank := line != ""
	isNotComment := !strings.HasPrefix(line, commentPrefix)
	return isNotBlank && isNotComment
}

// getNextContentRichLine will repeatedly ask the scanner for the next line
// until either: (1) the scanner is exhausted, or (2) a line meeting the content
// criteria is found.
func getNextContentRichLine(scanner *bufio.Scanner) (string, error) {
	var text string
	scanned := scanner.Scan()

	for scanned {
		text = scanner.Text()
		if isContentRich(text) {
			break
		}
		scanned = scanner.Scan()
	}

	if scanned {
		return text, nil
	} else {
		return "", io.EOF
	}
}
