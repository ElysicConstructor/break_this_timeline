package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const port = 5555

func initMSG() {
	fmt.Println("Willkommen zum Golang Peer-Chat :_> Dies ist das Xte mal einer Kausalen Scheiss-Schleife, Lets BREAK 1=1. diese redundanz ist immer der gleiche todes ausgangspunkt.!")
	fmt.Println("NJoin the BrainFuck <$.")
	fmt.Println("Gib 'help' ein, um Hilfe zu erhalten.")
}

func main() {

	initMSG()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Willst du Host oder Client sein? [H/C]: ")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(strings.ToUpper(mode))

	var targetHost string
	if mode == "C" {
		fmt.Print("Hostname oder IP des Servers: ")
		h, _ := reader.ReadString('\n')
		targetHost = strings.TrimSpace(h)
	}

	if mode == "H" {
		startServer()
	} else {
		// Versuche Verbindung
		conn, finalHost := connectToHost(targetHost)
		if conn != nil {
			fmt.Println("Verbunden mit:", finalHost)
			startClient(conn)
		} else {
			fmt.Println("Keine Verbindung möglich")
		}
	}
}

// ---------------------- SERVER ----------------------
func startServer() {
	fmt.Println("Starte Host auf Port", port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Fehler beim Starten:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Host läuft, warte auf Clients...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept Error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println("Neuer Client verbunden:", addr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("[%s] %s\n", addr, text)
	}
}

// ---------------------- CLIENT ----------------------
func connectToHost(hostname string) (net.Conn, string) {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("Hostname konnte nicht aufgelöst werden:", err)
		return nil, ""
	}

	for _, ip := range ips {
		address := fmt.Sprintf("%s:%d", ip.String(), port)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			return conn, ip.String()
		}
	}
	return nil, ""
}

func startClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Gib Nachrichten ein (Enter zum Senden, Ctrl+C zum Beenden):")
	scanner := bufio.NewScanner(os.Stdin)
	go receiveMessages(conn)

	for scanner.Scan() {
		text := scanner.Text()
		conn.Write([]byte(text + "\n"))
	}
}

func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("[Remote] " + scanner.Text())
	}
}
