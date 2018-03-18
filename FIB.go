// ROSALIND Problem 4

// In 1202, Leonardo of Pisa (commonly known as Fibonacci) considered a mathematical exercise regarding the reproduction of a population of rabbits. He made the following simplifying assumptions about the population:
// 1. The population begins in the first month with a pair of newborn rabbits.
// 2. Rabbits reach reproductive age after one month.
// 3. In any given month, every rabbit of reproductive age mates with another rabbit of reproductive age.
// 4. Exactly one month after two rabbits mate, they produce one male and one female rabbit.
// 5. Rabbits never die or stop reproducing.

// Fibonacci's exercise was to calculate how many pairs of rabbits would remain in one year.
// We can see that in the second month, the first pair of rabbits reach reproductive age and mate.
// In the third month, another pair of rabbits is born, and we have two rabbit pairs; our first pair of rabbits mates again.
// In the fourth month, another pair of rabbits is born to the original pair, while the second pair reach maturity and mate (with three total pairs), etc.

// Now, a sequence is an ordered collection of objects (usually numbers), which are allowed to repeat. Sequences can be finite or infinite.
// A recurrence relation is a way of defining the terms of a sequence with respect to the values of previous terms.
// In the case of Fibonacci's rabbits from the introduction, any given month will contain the rabbits that were alive the previous month, plus any new offspring.
// A key observation is that the number of offspring in any month is equal to the number of rabbits that were alive two months prior.
// As a result, if Fn represents the number of rabbit pairs alive after the n-th month, then we obtain the Fibonacci sequence having terms Fn that are defined by the recurrence relation Fn=Fn−1+Fn−2 (with F1=F2=1 to initiate the sequence).
// Note: n and n-1 and n-2 are meant to be subscripts.

// When finding the n-th term of a sequence defined by a recurrence relation, we can simply use the recurrence relation to generate terms for progressively larger values of n.
// This problem introduces us to the computational technique of dynamic programming, which successively builds up solutions by using the answers to smaller cases.

// Given: Positive integers n≤40 and k≤5.
// Return: The total number of rabbit pairs that will be present after n months if we begin with 1 pair and in each generation, every pair of reproduction-age rabbits produces a litter of k rabbit pairs (instead of only 1 pair).
// Sample Dataset: 5, 3
// Sample Output: 19

// Since each generation, every pair of reproduction-age rabbits produces k rabbit offspring, our formula changes to:
// Fn = Fn-1 + k*Fn-2 (where F1=F2=1)

package main

import "fmt"

func rabbitFibonacci() func() int64 { // need to use int64, not just int, because otherwise it will fail for big numbers of N
	F1, F2 := int64(1), int64(1) // F1=F2=1
	k := int64(3)                // set k=4
	F3 := F2 + k*F1

	return func() int64 {
		ans := F2
		F2, F3 = F3, F3+k*F2
		return ans
	}
}

func main() {
	f := rabbitFibonacci()
	for i := 0; i < 31; i++ {
		fmt.Println(f()) //this will print out F2 to F40
	}
}

// Note: a different way to do it (which can calculate the ans for very big values of N or k, e.g. N=1000 and more):
// package main

// import (
// 	"fmt"
// 	"math/big"
// )

// func rabbit_fibonacci() func() string {
// 	F1, F2 := big.NewInt(1), big.NewInt(1)
// 	k := big.NewInt(3)
// 	F3 := &big.Int{}
// 	F3.Mul(k, F1)
// 	F3.Add(F3, F2)

// 	tmp := &big.Int{}
// 	return func() string {
// 		ans := F2.String()

// 		//F2, F3 = F3, F3+k*F2
// 		tmp.Mul(k, F2)
// 		tmp.Add(F3, tmp)
// 		F2, F3, tmp = F3, tmp, F2

// 		return ans
// 	}
// }

// func main() {
// 	f := rabbit_fibonacci()
// 	for i := 0; i < 40; i++ {
// 		fmt.Println(f())
// 	}
// }
