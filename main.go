package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func scanPortsConcurrentlyWithSemaphore(hostname string, start, end int, maxRoutines int) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, maxRoutines) // Limit to maxRoutines at the same time

	fmt.Printf("Scanning ports %d to %d on %s...\n", start, end, hostname)
	for port := start; port <= end; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()

			semaphore <- struct{}{}        // Add a routine to the semaphore
			defer func() { <-semaphore }() // Remove the routine from the semaphore

			if scanPort("tcp", hostname, port) {
				fmt.Printf("Port %d is open\n", port)
			}
		}(port)
	}

	wg.Wait()
	fmt.Println("Scan complete.")
}

func main() {
	var hostname string
	var startPort, endPort, maxRoutines int

	fmt.Println("Welcome to PortHawk - Enhanced Port Scanner")
	fmt.Print("Enter IP address or hostname to scan: ")
	fmt.Scanln(&hostname)

	fmt.Print("Enter the start port: ")
	fmt.Scanln(&startPort)

	fmt.Print("Enter the end port: ")
	fmt.Scanln(&endPort)

	if startPort < 1 || endPort > 65535 || startPort > endPort {
		fmt.Println("Invalid port range. Please enter valid ports between 1 and 65535.")
		return
	}

	fmt.Print("Enter the maximum number of concurrent routines (e.g., 100): ")
	fmt.Scanln(&maxRoutines)

	if maxRoutines < 1 {
		fmt.Println("Invalid number of routines. Must be at least 1.")
		return
	}

	scanPortsConcurrentlyWithSemaphore(hostname, startPort, endPort, maxRoutines)
}
