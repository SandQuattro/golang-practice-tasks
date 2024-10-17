package main

import (
	"errors"
	"fmt"
	"github.com/rednafi/circuit-breaker/cb"
	"time"
)

func main() {
	// Initialize circuit breaker with:
	// - 3 failures threshold
	// - 5 seconds recovery time
	// - 2 requests allowed in half-open state
	// - 2 seconds request timeout
	circuitBreaker := cb.NewCircuitBreaker(3, 5*time.Second, 2, 2*time.Second)

	// Simulating a successful service request
	successFn := func() (any, error) {
		return "Success!", nil
	}

	result, err := circuitBreaker.Call(successFn)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Request succeeded: %v\n", result)
	}

	// Simulating a failed service request
	failFn := func() (any, error) {
		return nil, errors.New("service failure")
	}

	result, err = circuitBreaker.Call(failFn)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	} else {
		fmt.Printf("Request succeeded: %v\n", result)
	}
}
