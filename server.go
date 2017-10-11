package main

import "fmt"
import "net"

func runServer(sourceHost string, sourcePort int, targetHost string, targetPort int) {

	source := net.UDPAddr{IP: net.ParseIP(sourceHost), Port: sourcePort}
	sourceConn, err := net.ListenUDP("udp", &source)
	if err != nil {
		fmt.Printf("Could not establish connection on %v %v", source, err)
	}
	fmt.Printf("Starting UDP proxy from %v:%v -> %v:%v", sourceHost, sourcePort, targetHost, targetPort)
	for {
		buffer := make([]byte, 10240)
		n, addr, err := sourceConn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Printf("Could not read incoming packet from %v", addr)
		}

		go forward(targetHost, targetPort, buffer[:n])

	}

}

func forward(targetHost string, targetPort int, data []byte) {
	target := net.UDPAddr{IP: net.ParseIP(targetHost), Port: targetPort}
	targetConn, err := net.DialUDP("udp", nil, &target)

	if err != nil {
		fmt.Printf("Could not establish connection on %v:%v %v", targetHost, targetPort, err)
	}

	_, err = targetConn.Write(data)
	if err != nil {
		fmt.Printf("Could not forward the following data %v", string(data))
	}
}
