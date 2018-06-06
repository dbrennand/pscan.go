package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	connect(addr_input())
}

func addr_input() string {
	// Initiating Reader object to read user input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Example: www.hackthissite.org\n")
	fmt.Print("Server Address or IP: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read string: ", err)
	}
	// Done to remove the '\n' from str input. Otherwise, net.Dial will fail.
	addr := strings.Replace(input, "\n", "", -1)
	return addr
}

func connect(addr string) {
	// Slice holding commonly used ports.
	// http://packetlife.net/media/library/23/common-ports.pdf
	common_ports := [11]int{20, 21, 22, 23, 25, 53, 67, 68, 80, 110, 443}
	// Loops trying to connect to host at each port from slice above.
	for _, port := range common_ports {
		concat_addr := fmt.Sprintf("%v%v%v", addr, ":", port)
		conn, err := net.DialTimeout("tcp", concat_addr, 1*time.Second)
		// If connection =! nil: Tell user port is open.
		if conn != nil {
			fmt.Printf("Connection Successful - Port %v is open\n", port)
			// If err != nil: Tell user port is closed.
		} else if err != nil {
			fmt.Printf("Connection Unsuccessful - Port %v is closed\n", port)
		}
	}
}
