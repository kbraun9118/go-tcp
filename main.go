package main

import (
	"log"

	"github.com/songgao/water"
)

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}

	config.Name = "tun0"

	ifce, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1500)

	for {

		n, err := ifce.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Read: %d bytes\n", n)
		log.Printf("Buffer: %v\n", buffer)
	}
}
