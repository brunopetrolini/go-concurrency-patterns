package main

import "fmt"

func main() {
	jobs := make(chan int, 45)
	results := make(chan int, 45)

	go worker(jobs, results)
	go worker(jobs, results)

	for counter := 0; counter < 45; counter++ {
		jobs <- counter
	}
	close(jobs)

	for counter := 0; counter < 45; counter++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for number := range jobs {
		results <- fibonacci(number)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}
