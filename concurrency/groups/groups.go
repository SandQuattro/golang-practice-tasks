package main

import (
	"fmt"
	. "github.com/klauspost/cpuid/v2"
	"strings"
	"sync"
)

var links = []string{"1", "2", "3"}

func main() {
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	fmt.Println("Features:", strings.Join(CPU.FeatureSet(), ","))
	fmt.Println("Cacheline bytes:", CPU.CacheLine)
	fmt.Println("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", CPU.Cache.L1I, "bytes")
	fmt.Println("L2 Cache:", CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", CPU.Hz, "hz")

	counter := int64(0)
	linkCh := make(chan string, len(links))
	wg := sync.WaitGroup{}

	for i := range links {
		linkCh <- links[i]
	}
	close(linkCh)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range linkCh {
				fmt.Printf("link: %s", j)
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("counter: %d", counter)
}
