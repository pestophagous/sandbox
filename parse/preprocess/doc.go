// Preprocess is a small package that facilitates the shared commenting rules for
// both plain-text data formats (tab delimited and jsonl).
// With both formats it is safe and desireable to:
// - trim leading and trailing whitespace from each line
// - throw out blank lines
// - throw out lines that begin '//' (the comment marker)
