package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	addr    string
	port    = "80"
	list    int
	wait    int
	timeout int
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

	pass := 0
	var min float32 = float32(timeout) * 1e3
	var max float32 = 0
	var sum float32 = 0
	for i := 1; i <= list; i++ {
		p, t := ping(h, i)
		if t < min {
			min = t
		} else if t > max {
			max = t
		}
		sum += t
		pass += p
		time.Sleep(time.Duration(wait) * time.Second)
	}
	fmt.Println("Sent=", list, ", Successful=", pass, ", Failed=", list-pass, "(", (list-pass)/list, "% Fail)")
	fmt.Println("Minimum=", min, "ms, Maximum=", max, "ms, Average=", sum/float32(pass), "ms")
}

func parse() {
	h := flag.Bool("h", false, "Show Help")
	h2 := flag.Bool("help", false, "Show Help")
	flag.IntVar(&wait, "w", 1, "")
	flag.IntVar(&list, "l", 3, "")
	flag.IntVar(&timeout, "t", 2, "")
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
	fmt.Println("TCP Ping v0.1.1")
	fmt.Println("https://github.com/rehtt/tcping")
	fmt.Println("Use: tcping [-w] [-l] [-t] <IP address / Host> [Port (default: 80)]")
	fmt.Println("Must fill in IP address or Host.")
	fmt.Println("You can choose to fill in the port, port default 80.")
	fmt.Println("-w 5\t: ping every 5 seconds, default 1")
	fmt.Println("-l 5\t: send 5 pings, default 3")
	fmt.Println("-t 5\t: timeout 5 seconds, default 2")
	fmt.Println("eg: tcping google.com")
	fmt.Println("eg: tcping google.com 443")
	fmt.Println("eg: tcping -w 10 -l 6 -t 3 google.com 443")
	fmt.Println()
	os.Exit(0)
}

func ping(h *net.TCPAddr, i int) (pass int, pingt float32) {
	open := "is open"
	t := time.Now()
	c, err := net.DialTimeout("tcp", h.String(), time.Duration(timeout)*time.Second)
	pingt = float32(time.Now().UnixNano()-t.UnixNano()) / 1e6
	if err != nil {
		open = "no response"
		pass = 0
	} else {
		c.Close()
		pass = 1
	}
	fmt.Println(i, "- Ping", h.String(), "/tcp - Port", open, "- time=", pingt, "ms")
	return pass, pingt
}