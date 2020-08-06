package globalip

import (
	"fmt"
	"log"

	webrtc "github.com/pion/webrtc/v2"
)

func GetIPaddr() string {
	ch := make(chan string)
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	// generate a new connection
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
		} else {
			switch c.Typ {
			case webrtc.ICECandidateTypeHost:
				log.Println("Local IP Address:", c.Address)
			case webrtc.ICECandidateTypeSrflx:
				log.Println("Public IP Address:", c.Address)
				ch <- c.Address
				close(ch)
			}
		}
	})

	if _, err := peerConnection.CreateDataChannel("", nil); err != nil {
		log.Fatal(err)
	}

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Fatal(err)
	}

	if err = peerConnection.SetLocalDescription(offer); err != nil {
		log.Fatal(err)
	}

	// block forever
	select {
	case addr := <-ch:
		return addr
	}
	fmt.Println("done")
	return "error"
}
