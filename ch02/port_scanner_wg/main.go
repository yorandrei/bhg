package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

var start_port int
var end_port int
var ip string
var ch_size int

func init() {
	flag.IntVar(&start_port, "s", 1, "Specifies the port from which the scan begins")
	flag.IntVar(&end_port, "e", 65535, "Specifies the port at which the scan ends")
	flag.StringVar(&ip, "i", "127.0.0.1", "Specifies the IP address of the scan target")
	flag.IntVar(&ch_size, "c", 100, "Ports channel size")
}

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", ip, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	flag.Parse()
	ports := make(chan int, ch_size)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := start_port; i < end_port; i++ {
			ports <- i
		}
	}()

	for i := start_port; i < end_port; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
