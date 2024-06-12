// mcherald
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

const broadcastAddr = "224.0.2.60:4445"

var motdPortRegex = regexp.MustCompile(`(.*):(\d+)$`)

var (
	verboseFlag   = flag.Bool("v", false, "Verbose output")
	delayTimeFlag = flag.Duration("t", 1500*time.Millisecond, "Time between broadcasts")
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Error: No motd:port arguments specified.")
		return
	}
	// Collect UDP messages to send
	var motdPorts []string
	usedPorts := make(map[int64]bool)
	for _, arg := range flag.Args() {
		match := motdPortRegex.FindStringSubmatch(arg)
		if match == nil {
			fmt.Fprintln(os.Stderr, "Error: Argument not in the form motd:port: ", arg)
			return
		}
		portNum, _ := strconv.ParseInt(match[2], 10, 64)
		if portNum < 1 {
			fmt.Fprintln(os.Stderr, "Error: Port not > 0 in argument ", arg)
			return
		}
		if usedPorts[portNum] {
			fmt.Fprintln(os.Stderr, "Error: Port", portNum, "specified more than once!")
			return
		}
		usedPorts[portNum] = true
		newMsg := fmt.Sprintf("[MOTD]%s[/MOTD][AD]%d[/AD]", match[1], portNum)
		motdPorts = append(motdPorts, newMsg)
	}

	// Initialize network
	minecraftAddrPort, err := net.ResolveUDPAddr("udp", broadcastAddr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Couldn't resolve address ", broadcastAddr, ": ", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, minecraftAddrPort)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Couldn't open connection: ", err)
		return
	}
	// Broadcast forever
	if *verboseFlag {
		fmt.Println("Starting main loop.")
	}
	for {
		for _, msg := range motdPorts {
			if *verboseFlag {
				fmt.Println("Broadcasting", msg)
			}
			conn.Write([]byte(msg))
		}
		time.Sleep(*delayTimeFlag)
	}
}
