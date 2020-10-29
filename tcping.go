package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	addr string
	port = "80"
	list int
	wait int
)

func main() {
	parse()
	h, err := net.ResolveTCPAddr("tcp", addr+":"+port)
	if err != nil {
		fmt.Println()
		fmt.Println("DNS: Could not find host -", addr+":"+port)
		fmt.Println("'tcping -h' View help")
		fmt.Println()
		return
	}

	for i := 1; i <= list; i++ {
		open := "is open"
		t := time.Now()
		c, err := net.DialTimeout("tcp", h.String(), 2*time.Second)
		pingt := float32(time.Now().UnixNano()-t.UnixNano()) / 1e6
		if err != nil {
			open = "no response"
		} else {
			c.Close()
		}
		fmt.Println(i, "- Ping", h.String(), "/tcp - Port", open, "- time=", pingt, "ms")
		time.Sleep(time.Duration(wait) * time.Second)
	}
}

func parse() {
	h := flag.Bool("h", false, "Show Help")
	h2 := flag.Bool("help", false, "Show Help")
	flag.IntVar(&wait, "w", 1, "wait")
	flag.IntVar(&list, "l", 3, "list")
	flag.Parse()
	if *h || *h2 {
		help()
	}
	addr = flag.Arg(0)
	if addr == "" {
		help()
	}
	if len(flag.Args()) == 2 {
		port = flag.Arg(1)
	}
}

func help() {
	fmt.Println()
	fmt.Println("TCP Ping v0.1")
	fmt.Println("https://github.com/rehtt/tcping")
	fmt.Println("Use: tcping [-w] [-l] <IP address / Host> [Port (default: 80)]")
	fmt.Println("Must fill in IP address or Host.")
	fmt.Println("You can choose to fill in the port, port default 80.")
	fmt.Println("-w 5\t: ping every 5 seconds, default 1")
	fmt.Println("-l 5\t: send 5 pings, default 3")
	fmt.Println("eg: tcping google.com")
	fmt.Println("eg: tcping google.com 443")
	fmt.Println("eg: tcping -w 10 -l 6 google.com 443")
	fmt.Println()
	os.Exit(0)
}
