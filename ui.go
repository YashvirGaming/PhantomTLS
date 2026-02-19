package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strings"
	"github.com/fatih/color"
)

const TITLE = "PHANTOMTLS"
const SUBTITLE = "Made with <3 ♥ by Yashvir Gaming"

func neonPulse(text string) {
	for i := 0; i < 2; i++ {
		color.HiCyan(text)
		time.Sleep(200 * time.Millisecond)
		color.HiMagenta(text)
		time.Sleep(200 * time.Millisecond)
	}
}

func showBanner() {
	fmt.Println()
	neonPulse("████████  PHANTOMTLS  ████████")
	color.New(color.FgHiYellow, color.Bold).Printf(">> %s <<\n", TITLE)
	color.HiMagenta(SUBTITLE)
	fmt.Println()
}

func showListenerInstructions() {

	color.HiGreen("\n[+] make requests to localhost and the port you've chosen")
	fmt.Println("using either the POST or GET method, respectively.")

	color.HiGreen("[+] make sure to include the below headers accompanied by their values:\n")

	color.HiCyan("-- [required]")
	fmt.Println("x-url: https://example.com")
	fmt.Println("x-proxy: http://user:pass@ip:port")
	fmt.Println("x-identifier: chrome | firefox | safari_ios | safari | okhttp")
	fmt.Println("x-session-id: a random string to keep track of request sessions IE: PhantomTLS\n")

	color.HiCyan("-- [optional]")
	fmt.Println("host: example.com")
	fmt.Println("content-encoding: gzip\n")

	color.HiYellow("note: the client assumes both content-type and content-length headers are being passed.\n")

	color.HiMagenta("Server is running on http://localhost:2024\n")
}

func showMenu() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("[1] Reverse proxy Mode [Low Speed]")
	fmt.Println("[2] Multithread Mode [20x Fastest Speed]")
	fmt.Println("[3] Proxy Mode [Medium Speed")
	fmt.Println("[4] TLS Handshake Emulator [Faster Speed]")
	fmt.Println("[5] Exit")
	fmt.Print("\nSelect option: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		showListenerInstructions()
		startServer()
	case "2":
		startWorkerMode()
		waitForEnter()
	case "3":
		startProxyListMode()
		waitForEnter()
	case "4":
		startTLSMode()
		waitForEnter()
	case "5":
		os.Exit(0)
	default:
		fmt.Println("Invalid option.")
		waitForEnter()
	}
}


func waitForEnter() {
	fmt.Print("\nPress ENTER to return to menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println()
	showMenu()
}
