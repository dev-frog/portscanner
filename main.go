package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}

func main() {
	ports := make(chan int, 100)
	result := make(chan int)

	var openPorts []int

	for i := 1; i <= cap(ports); i++ {
		go worker(ports, result)
	}

	go func() {
		for i := 1; 1 <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-result
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(result)

	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 1024; i++ {
// 		wg.Add(1)
// 		go func(j int) {
// 			defer wg.Done()
// 			address := fmt.Sprintf("scanme.nmap.org:%d", j)
// 			conn, err := net.Dial("tcp", address)

// 			if err != nil {
// 				return
// 			}
// 			conn.Close()
// 			fmt.Printf("%d open port\n", j)
// 		}(i)
// 	}
// 	wg.Wait()
// }
