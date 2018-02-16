package main

import (
	"log"
	"time"

	"github.com/nictuku/dht"
)

func main() {
	// START OMIT
	node, _ := dht.New(nil) // TODO: error check
	node.Start()            // TODO: error check

	dataHash, _ := dht.DecodeInfoHash(`13c6d46f99e6e849176485ba5c8d5ab460f9eccc`)
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for {
			log.Println("Requesting peers")
			node.PeersRequest(string(dataHash), true)
			<-ticker.C
		}
	}()
	for results := range node.PeersRequestResults {
		for hash, peers := range results {
			if hash != dataHash {
				continue
			}
			for _, peer := range peers {
				log.Printf("Found peer: %s\n", dht.DecodePeerAddress(peer))
			}
		}
	}
	// END OMIT
}
