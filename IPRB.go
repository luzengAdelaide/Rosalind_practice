// Rosalind problem 9
// Introduction to Mendelian Inheritance
// Probability is the mathematical study of randomly occurring phenomena.
// We will model such a phenomenon with a random variable,
// which is simply a variable that can take a number of different distinct outcomes depending on the result of an underlying random process.

// For example, say that we have a bag containing 3 red balls and 2 blue balls.
// If we let X represent the random variable corresponding to the color of a drawn ball,
// then the probability of each of the two outcomes is given by Pr(X=red)=35 and Pr(X=blue)=25.

// Given: Three positive integers k, m, and n, representing a population containing k+m+n organisms:
// k individuals are homozygous dominant for a factor, m are heterozygous, and n are homozygous recessive.
//
// Return: The probability that two randomly selected mating organisms will produce an individual possessing
// a dominant allele (and thus displaying the dominant phenotype). Assume that any two organisms can mate.

package main

import (
	"fmt"
)

// Method 1:

/*func MendleFirstLaw(dom float64, hetero float64, rec float64) float64 {

	// probability of each mating scenario that could produce a dominant phenotype, plus probability of phenotype in mating

	dom2 := 1.0 + (dom-2)*2.0
	domhetero := dom * hetero
	domrec := dom * rec
	hetero2 := 1.0 + (hetero-2)*2.0
	heterorec := hetero * rec
	rec2 := 1.0 + (rec-2)*2.0

	//pheterorec := phetero * prec * 0.5 * 2

	// summing across probabilities (1-probability of a recessive phenotype)

	dom_total := dom2 + domhetero + domrec + 0.75*hetero2 + 0.5*heterorec
	total := dom2 + domhetero + domrec + hetero2 + heterorec + rec2
	return dom_total / total
}*/

// Method2:

func MendleFirstLaw2(dom float64, hetero float64, rec float64) float64 {

	// probability of each mating scenario that could produce a dominant phenotype, plus probability of phenotype in mating
	total := dom + hetero + rec
	rec2 := (rec / total) * ((rec - 1) / (total - 1))
	hetero2 := (hetero / total) * ((hetero - 1) / (total - 1))
	rechetero := (rec/total)*(hetero/(total-1)) + (hetero/total)*(rec/(total-1))

	rec_total := rec2 + 0.25*hetero2 + 0.5*rechetero
	return 1 - rec_total
	// summing across probabilities (1-probability of a recessive phenotype)

}

func main() {
	fmt.Print(MendleFirstLaw2(22, 20, 28))
}
