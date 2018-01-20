package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"os"
	"strings"
)

//ui holds the userinput
var ui = ""

func main() {
	argCount := len(os.Args[1:])
	if argCount != 1 {
		fmt.Printf("Usage: %s <serial-port>\ne.g.:\n\t%s com5 (on windows)\n\t%s /dev/ttyUSB0 (on linux)\n\t%s /dev/cu.usbmodem1411 (on osx)\n", os.Args[0], os.Args[0], os.Args[0], os.Args[0])
		os.Exit(0)
	}
	c := &serial.Config{Name: os.Args[1], Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println("chamgo - ready to talk to the chameleon")
	for {
		ui = getUserInput()
		if ui == "exit" || ui == "quit" || ui == "q" {
			os.Exit(0)
		}

		rb := sendCmd(ui, s)
		fmt.Print(string(rb))
	}
}

/**
* sending cmds to serial
* returns {bytes} responde from device
**/
func sendCmd(cmd string, port *serial.Port) (returnmessage []byte) {
	cmd += "\r"
	n, err := port.Write([]byte(cmd))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 256)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	return buf[0:n]
}

/**
* for receiving user input
* returns {string} with no newline (|n)
**/
func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	return strings.Replace(string(line), "\n", "", -1)
}
