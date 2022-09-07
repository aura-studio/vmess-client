package main

import (
	"log"

	"github.com/aura-studio/vmess-client/vmesshub"
)

const (
	uuid            = "832374fd-a2f6-4098-8dfb-0d26d9032d40"
	socks5ListenUri = "0.0.0.0:1080"
	remoteHost      = "23.83.224.64:21537"
	alterId         = 0
	securityType    = "chacha20-poly1305"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Ltime | log.Lshortfile)

	log.Println("Local Socks5 :", socks5ListenUri)

	vh, err := vmesshub.CreateVmessHub(socks5ListenUri, remoteHost, uuid, securityType, alterId)

	if err != nil {
		log.Println(err)
		return
	}

	vh.StartSocks5Listen()
}
