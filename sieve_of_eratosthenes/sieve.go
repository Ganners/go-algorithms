// This is an implementation of the prime sieve, but rather than returning the
// first N prime numbers, it will give you the prime numbers up to N. More
// useful for my purposes
package sieve

// Aggregates and returns the sieve
func Aggregate(primesChan chan int, doneChan chan struct{}) []int {

	// Syncronise and aggregate
	results := make([]int, 0)

aggregate:
	for {
		select {
		case prime := <-primesChan:
			results = append(results, prime)
		case <-doneChan:
			break aggregate
		}
	}

	return results
}

// SieveOfEratosthenes will execute up to a given number and send back
// values along an integer channel
func SieveOfEratosthenes(upTo int) (chan int, chan struct{}) {

	if upTo < 2 {
		return nil, nil
	}

	primesChan := make(chan int)
	doneChan := make(chan struct{})
	feed := feedNumbers(upTo)

	go func() {
		for {
			prime := <-feed

			// If it's prime, pass done
			if prime == 0 {
				doneChan <- struct{}{}
			}

			// Pass prime to primesChan
			primesChan <- prime

			// Filter and swap this feed to the out feed from this
			outFeed := shake(feed, prime)
			feed = outFeed
		}
	}()

	return primesChan, doneChan
}

// Shake the sieve (filter), send back anything received that is prime
func shake(feed <-chan int, prime int) chan int {

	primesChan := make(chan int)

	go func() {
		for {

			number := <-feed

			if number == 0 {
				primesChan <- 0
				continue
			}

			if number%prime != 0 {
				primesChan <- number
			}
		}
	}()

	return primesChan
}

// Creates a number feed and passes numbers along it
func feedNumbers(upTo int) <-chan int {

	feed := make(chan int)

	go func() {
		for i := 2; i <= upTo; i++ {

			feed <- i

			if i >= upTo {
				// Send a 0 for termination
				feed <- 0
			}
		}
	}()

	return feed
}
