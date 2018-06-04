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
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read string: ", err)
	}
	// Done to remove the '\n' from str input. Otherwise, net.Dial will fail.
	addr_port := strings.Replace(input, "\n", "", -1)
	return addr_port
}

func connect(addr_port string) {
	// addr_port is passed into .Dial function.
	conn, err := net.Dial("tcp", addr_port)
	// Split str at ':' to get the specific port number.
	port := strings.Split(addr_port, ":")
	// If Err tell user connection failed and presented err.
	if err != nil {
		fmt.Printf("Connection Unsuccessful - Port %v is closed\n", port[1])
		fmt.Println("Error: ", err)
		// else if the connection is successful, tell user.
	} else if conn != nil {
		fmt.Printf("Connection Successful - Port %v is open\n", port[1])
	}
	// Close the connection.
	conn.Close()
}
