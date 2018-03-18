package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

// define a new "seq" type which is a struct containing fields
type seq struct {
	name string
	seq  []byte
	gc   float64
}

// make a new type, like a collection of seq
type seqs []seq

// make a new function to calculate the GC content of FASTA seqs
// i.e. the percentage of symbols in the string that are 'C' or 'G'
func gc(s []byte) float64 {
	var a [128]byte
	for _, b := range s {
		a[b]++
	}
	return ((float64(a['C']) + float64(a['G'])) / float64(len(s))) * 100
}

// We're going to need to sort the seqs in order from highest to lowest GC content
// This defines the sort function
func (s seqs) Len() int           { return len(s) }
func (s seqs) Less(i, j int) bool { return s[i].gc > s[j].gc }
func (s seqs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {

	// Feed in a FASTA file from the set directory
	inputFile, err := os.Open("rosalind_gc.txt")
	if err != nil {
		log.Fatalf("reading file: %v", err)
	}
	defer inputFile.Close()

	var ss seqs
	sc := bufio.NewScanner(inputFile)
	for sc.Scan() {

		// Trim any white space (e.g. blank lines in the FASTA file)
		l := bytes.TrimSpace(sc.Bytes())
		if len(l) == 0 {
			continue
		}
		// Now this is where things start getting a bit convoluted
		// We don't know how long each sequence is - but we can't calculate the GC content until we get to the end of the sequence
		// If we hit another '>' sign, that signals the end of the previous seq
		// So at the point, we know the length of the seq, and can work out the GC content
		// That's what the below code is doing:
		// It's basically saying, if we have a '>' (and it's not a blank seq), then calculate the GC content of the previous seq
		// l[1:] read name from begin to end
		// l... read sequence from begin to end
		switch l[0] {
		case '>':
			if len(ss) != 0 {
				ss[len(ss)-1].gc = gc(ss[len(ss)-1].seq)
			}
			ss = append(ss, seq{name: string(l[1:])})
		default:
			ss[len(ss)-1].seq = append(ss[len(ss)-1].seq, l...)
		}
	}
	if sc.Err() != nil {
		log.Fatalf("failed scanning: %v", err)
	}
	// if the length of the sequence is 0, there is no sequence
	if len(ss) == 0 {
		log.Println("no sequence")
		os.Exit(0)
	}

	// The above worked for every seq that had another '>' below it
	// But what about the last sequence in the file?
	// This line calculates the GC content for that last seq, outside of the loop
	ss[len(ss)-1].gc = gc(ss[len(ss)-1].seq)

	// Need to sort the seqs in order from highest to lowest GC content
	sort.Sort(ss)

	// Print out the seq header and GC content
	for i := 0; i < len(ss); i++ {
		fmt.Printf("%s %v\n", ss[i].name, ss[i].gc)
	}
}
