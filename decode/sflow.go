package decode

import (
	"github.com/PreetamJinka/cistern/net/sflow"
)

type SflowDecoder struct {
	inbound  <-chan []byte
	outbound chan sflow.Datagram
}

func NewSflowDecoder(inbound <-chan []byte, bufferLength ...int) *SflowDecoder {
	bufLen := 0

	if len(bufferLength) > 0 {
		bufLen = bufferLength[0]
	}

	return &SflowDecoder{
		inbound:  inbound,
		outbound: make(chan sflow.Datagram, bufLen),
	}
}

func (d *SflowDecoder) Outbound() chan sflow.Datagram {
	return d.outbound
}

func (d *SflowDecoder) Run() {
	go func() {
		for buf := range d.inbound {
			d.outbound <- sflow.Decode(buf)
		}
	}()
}