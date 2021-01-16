package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram = make(map[rune]int)

	for _, base := range "ATCG" {
		h[base] = 0
	}

	for _, nucleotide := range d {
		if !strings.Contains("ACTG", string(nucleotide)) {
			return nil, errors.New("Unknown nucleotide found")
		}

		h[nucleotide]++
	}

	return h, nil
}
