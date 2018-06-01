package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	connect(addr_port_input())
}

func addr_port_input() string {
	// Initiating Reader object to read user input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Example: www.hackthissite.org:80\n")
	fmt.Print("Server Address & Port: ")
	addr_port, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read string: ", err)
	}
	return addr_port
}

func connect(addr_port string) {
	// I believe the issue is with .Dial() or im doing something obviously wrong
	// and just cant see it.
	conn, err := net.Dial("tcp", addr_port)
	port := strings.Split(addr_port, ":")
	if err != nil {
		fmt.Println("Error occured when connecting:", err)
		fmt.Printf("Port %v is closed", port[1])
	} else {
		fmt.Printf("Port %v is open", port[1])
	}
	conn.Close()
}
