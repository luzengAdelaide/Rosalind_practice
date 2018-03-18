// 1. read sequences files (fasta) format
// 2. separate data by each column
// 3. calculate each base ACGT
// 4. sort the frequency exist of there bases
// 5. print out the base with most happened in each column and then print out how they looks like in each column

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// definde a new "seq" type which is a struct containing fields
type seq struct {
	name  string
	seq   []byte
	count int
}

// make a new type, like a collection of seq
type seqs []seq

// make a new function to calculate the seqs
func count(s []byte) float64 {
	var a [128]byte
	for _, b := range s {
		a[b]++
	}
	return (float64(a[b]))
}

// sort the seqs in order from highest to lowest with each column
func main() {

	// Feed in a FASTA file from the set directory
	inputFile, err := os.Open("rosalind_lcsm.txt")
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
		switch l[0] {
		case '>':
			if len(ss) != 0 {
				ss[len(ss)-1].count = count(ss[len(ss)-1].seq)
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

	ss[len(ss)-1].count = count(ss[len(ss)-1].seq)

	sort.Sort(ss)

	for i := 0; i < len(ss); i++ {
		fmt.Printf(("%s %v\n"), ss[i].string, ss[i].count)
	}
}
