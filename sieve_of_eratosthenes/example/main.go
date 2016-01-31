package main

import (
	"flag"
	"fmt"

	"github.com/Ganners/go-algorithms/sieve_of_eratosthenes"
)

type Config struct {
	UpTo int
}

func parseFlags() *Config {

	c := &Config{}

	flag.IntVar(
		&c.UpTo, "upto", 5,
		"The value up to which you would like to print the prime numbers")

	flag.Parse()

	return c
}

func main() {

	conf := parseFlags()
	fmt.Printf("Printing primes up to %d:\n", conf.UpTo)

	primes := sieve.Aggregate(sieve.SieveOfEratosthenes(conf.UpTo))
	for _, prime := range primes {
		fmt.Printf(" - %d\n", prime)
	}

}
