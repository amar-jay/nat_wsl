package portmap

import (
	"fmt"
	"io"
	"log"
	"net"
)

// PortMapping represents a mapping between an internal and external port.
type PortMapping struct {
	Protocol     string // only TCP supported for now
	Internalip   string // windows host
	Internalport int    // windows port
	Externalip   string // the wsl host
	Externalport int    // the wsl port
}

func (pm PortMapping) String() string {
	return fmt.Sprintf("%s:%d -> %s:%d", pm.Internalip, pm.Internalport, pm.Externalip, pm.Externalport)
}

/**
* This function is used to forward packets from one network interface to another.
* It is used to forward packets from the WSL2 VM to the host machine.
 */
func (pm *PortMapping) Start() {
	host := fmt.Sprintf("%s:%d", pm.Internalip, pm.Internalport)

	// listen for incoming packets
	listener, err := net.Listen(pm.Protocol, host)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
	defer listener.Close()

	log.Printf("Listening on port %s, forwarding to %s:%d", host, pm.Externalip, pm.Externalport)

	for {
		// accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// its a two way connection

		// send the connection to the handleConnection function
		go handleConnectiontoWsl(conn, pm.Protocol, pm.Externalip, pm.Externalport)
	}
}

// handleConnection forwards data between a source and destination connection.
func handleConnectiontoWsl(src net.Conn, protocol, wslIP string, wslPort int) {
	defer src.Close()

	// connect to the destination
	dst, err := net.Dial(protocol, fmt.Sprintf("%s:%d", wslIP, wslPort))
	if err != nil {
		log.Printf("Failed to connect to %s:%d: %v", wslIP, wslPort, err)
		return
	}
	defer dst.Close()

	// Copy data between source and destination
	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Printf("Error copying data from source to destination: %v", err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Printf("Error copying data from destination to source: %v", err)
	}
}
