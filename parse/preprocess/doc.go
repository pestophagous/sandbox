// Preprocess is a small package that facilitates the shared commenting rules
// for both of our plain-text data formats (tab-delimited, and jsonl).  With
// both formats it is safe and desireable to:
// - trim leading and trailing whitespace from each line
// - throw out blank lines
// - throw out lines that begin '//' (the comment marker)
// (Note that comments are all-or-nothing: either the whole line is a comment,
// or else no part of the line is a comment.)
package preprocess
