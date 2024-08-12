package nat

import (
	"flag"
	"fmt"
	"log"

	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"

	"github.com/amar-jay/nat_wsl/pkg/config"
)

func NAT(conf *config.Forwarding) {
	// Initialize the NFF-GO library
	config := flow.Config{}
	flow.CheckFatal(flow.SystemInit(&config))

	// Create a packet flow
	packetFlow, err := flow.SetReceiver(uint16(*listenPort))
	flow.CheckFatal(err)

	// Process packets
	flow.CheckFatal(flow.SetHandler(packetFlow, func(pkt *packet.Packet, context flow.UserContext) {
		handlePacket(pkt, *connectAddress, *connectPort)
	}, nil))

	// Send packets
	flow.CheckFatal(flow.SetSender(packetFlow, uint16(*connectPort)))

	// Start processing
	flow.CheckFatal(flow.SystemStart())
}

func handlePacket(pkt *packet.Packet, connectAddress string, connectPort int) {
	// This is a stub for handling packet forwarding
	// You can modify the packet here as needed
	if pkt == nil || pkt.Ether == nil {
		log.Println("Invalid packet received")
		return
	}

	// Example: Print packet info (for debugging purposes)
	fmt.Printf("Packet received: %+v\n", pkt)

	// Modify the packet to change the destination IP and port
	ipv4 := pkt.GetIPv4()
	if ipv4 != nil {
		ipv4.DstAddr = packet.SwapBytesUint32(packet.ParseIP4(connectAddress).To4())
	}

	udp, err := pkt.GetUDPForIPv4()
	if err == nil {
		udp.DstPort = packet.SwapBytesUint16(uint16(connectPort))
	}
}
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
)

func main() {
	// Parse command-line arguments
	listenPort := flag.Int("listenport", 0, "Port to listen on")
	connectPort := flag.Int("connectport", 0, "Port to connect to in WSL")
	connectAddress := flag.String("connectaddress", "", "WSL IP address to connect to")
	flag.Parse()

	if *listenPort == 0 || *connectPort == 0 || *connectAddress == "" {
		log.Fatal("All parameters (listenport, connectport, connectaddress) are required")
	}

	// Initialize the NFF-GO library
	config := flow.Config{}
	flow.CheckFatal(flow.SystemInit(&config))

	// Create a packet flow
	packetFlow, err := flow.SetReceiver(uint16(*listenPort))
	flow.CheckFatal(err)

	// Process packets
	flow.CheckFatal(flow.SetHandler(packetFlow, func(pkt *packet.Packet, context flow.UserContext) {
		handlePacket(pkt, *connectAddress, *connectPort)
	}, nil))

	// Send packets
	flow.CheckFatal(flow.SetSender(packetFlow, uint16(*connectPort)))

	// Start processing
	flow.CheckFatal(flow.SystemStart())
}

func handlePacket(pkt *packet.Packet, connectAddress string, connectPort int) {
	// This is a stub for handling packet forwarding
	// You can modify the packet here as needed
	if pkt == nil || pkt.Ether == nil {
		log.Println("Invalid packet received")
		return
	}

	// Example: Print packet info (for debugging purposes)
	fmt.Printf("Packet received: %+v\n", pkt)

	// Modify the packet to change the destination IP and port
	ipv4 := pkt.GetIPv4()
	if ipv4 != nil {
		ipv4.DstAddr = packet.SwapBytesUint32(packet.ParseIP4(connectAddress).To4())
	}

	udp, err := pkt.GetUDPForIPv4()
	if err == nil {
		udp.DstPort = packet.SwapBytesUint16(uint16(connectPort))
	}
}

