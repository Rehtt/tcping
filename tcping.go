package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	addr    string
	ports   [][]string
	list    int
	wait    int
	timeout int
)

func main() {
	parse()
	result := ""
	for _, vp := range ports {
		startPort, _ := strconv.Atoi(vp[0])
		endPort := startPort
		if len(vp) > 1 {
			endPort, _ = strconv.Atoi(vp[1])
		}

		for ; startPort <= endPort; startPort++ {
			h, err := net.ResolveTCPAddr("tcp", addr+":"+strconv.Itoa(startPort))
			if err != nil {
				fmt.Printf("\nDNS: Could not find host - %s : %d\n", addr, startPort)
				fmt.Println(`"tcping -h" View help`)
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
				}
				if t > max {
					max = t
				}
				sum += t
				pass += p
				time.Sleep(time.Duration(wait) * time.Second)
			}
			fmt.Printf("Sent= %d, Successful= %d, Failed= %d( %d%% Fail)\n", list, pass, list-pass, (list-pass)/list*100)
			pp := " no response\n"
			if pass > 0 {
				fmt.Printf("Minimum= %.2fms, Maximum= %.2fms, Average= %.2fms\n", min, max, sum/float32(pass))
				pp = " is Open\n"
			}
			fmt.Println()
			result += strconv.Itoa(startPort) + "/tcp" + pp
		}
	}
	fmt.Println(result)
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
	if len(flag.Args()) > 1 {
		for i := 1; i < len(flag.Args()); i++ {
			p := strings.Split(flag.Arg(i), "-")
			if len(p) > 2 {
				help()
			}
			var port []string
			for _, v := range p {
				if v != "" {
					port = append(port, v)
				}
			}
			if port == nil {
				help()
			}
			ports = append(ports, port)

		}

	} else {
		ports = append(ports, []string{"80"})
	}
}

func help() {
	fmt.Println()
	fmt.Println("TCP Ping v0.2.0")
	fmt.Println("https://github.com/rehtt/tcping")
	fmt.Println("Use: tcping [-w] [-l] [-t] <IP address / Host> [Port (default: 80)]")
	fmt.Println("Must fill in IP address or Host.")
	fmt.Println("You can choose to fill in the port, you can add multiple ports or use \"-\" to specify the range, port default 80.")
	fmt.Println("-w 5\t: ping every 5 seconds, default 1")
	fmt.Println("-l 5\t: send 5 pings, default 3")
	fmt.Println("-t 5\t: timeout 5 seconds, default 2")
	fmt.Println("eg: tcping google.com")
	fmt.Println("eg: tcping google.com 443")
	fmt.Println("eg: tcping google.com 80 443")
	fmt.Println("eg: tcping google.com 80-85 443-448")
	fmt.Println("eg: tcping -w 10 -l 6 -t 3 google.com 443")
	fmt.Println()
	os.Exit(0)
}

func ping(h *net.TCPAddr, i int) (pass int, pingt float32) {
	open := "is open"
	t := time.Now()
	c, err := net.DialTimeout("tcp", h.String(), time.Duration(timeout)*time.Second)
	pingt = float32(time.Now().Sub(t).Nanoseconds()) / 1e6
	if err != nil {
		open = "no response"
		pass = 0
	} else {
		c.Close()
		pass = 1
	}
	fmt.Printf("%d - Ping %s /tcp - Port %s - time= %.2fms\n", i, h.String(), open, pingt)
	return pass, pingt
}
