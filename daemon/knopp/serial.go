package main

import (
	"io"
	"strconv"

	"github.com/jacobsa/go-serial/serial"
)

const (
	baud     = uint(115200)
	dataBits = 8
	parity   = serial.PARITY_NONE
	stopBits = 1
)

type Event int

const (
	EventDown = iota
	EventUp
	EventLeft
	EventRight
)

type Knopp struct {
	Ch chan Event

	device string
	serial io.ReadWriteCloser

	curPos int
}

func NewKnopp(device string) (*Knopp, error) {
	k := Knopp{
		Ch:     make(chan Event),
		device: device,
	}

	err := k.initSerial()
	if err != nil {
		return nil, err
	}

	go k.handleEvents()
	return &k, nil
}

func (k *Knopp) initSerial() error {
	options := serial.OpenOptions{
		PortName:              k.device,
		BaudRate:              baud,
		DataBits:              dataBits,
		ParityMode:            parity,
		StopBits:              stopBits,
		MinimumReadSize:       1,
		InterCharacterTimeout: 100,
	}

	var err error
	k.serial, err = serial.Open(options)
	return err
}

func (k *Knopp) handleInput(s string) {
	switch s {
	case "U":
		k.Ch <- EventUp
	case "D":
		k.Ch <- EventDown

	default:
		np, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		if np < k.curPos {
			k.Ch <- EventLeft
		} else {
			k.Ch <- EventRight
		}
		k.curPos = np
	}
}

func (k *Knopp) handleEvents() {
	var b []byte
	buf := make([]byte, 8)

	for {
		n, err := k.serial.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			return
		}

		b = append(b, buf[:n]...)
		var s string
		var idx int
		for i, v := range b {
			if v == '\r' {
				continue
			}
			if v == '\n' {
				k.handleInput(s)

				s = ""
				idx = i
				continue
			}
			s += string(v)
		}

		b = b[idx+1:]
	}
}
