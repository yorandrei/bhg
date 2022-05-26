package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var start_port int
var end_port int
var ip string

func init() {
	flag.IntVar(&start_port, "s", 0, "Specifies the port from which the scan begins")
	flag.IntVar(&end_port, "e", 65535, "Specifies the port at which the scan ends")
	flag.StringVar(&ip, "i", "127.0.0.1", "Specifies the IP address of the scan target")
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup
	for i := start_port; i < end_port; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ip, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}

	wg.Wait()
}
