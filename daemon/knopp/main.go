package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bendahl/uinput"
)

func initVirtualDevices() (uinput.Dial, uinput.Keyboard) {
	dial, err := uinput.CreateDial("/dev/uinput", []byte("knopp"))
	if err != nil {
		log.Fatalf("Could not create virtual input device (/dev/uinput): %s", err)
	}

	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("knopp Keys"))
	if err != nil {
		log.Fatalf("Could not create virtual input device (/dev/uinput): %s", err)
	}

	return dial, keyboard
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: requires device name as parameter.")
		os.Exit(1)
	}

	mouse, keyboard := initVirtualDevices()
	defer mouse.Close()
	defer keyboard.Close()

	k, err := NewKnopp(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	for v := range k.Ch {
		switch v {
		case EventUp:
			fmt.Println("UP")
			err := keyboard.KeyUp(0x120)
			if err != nil {
				log.Fatal(err)
			}
		case EventDown:
			fmt.Println("DOWN")
			err := keyboard.KeyDown(0x120)
			if err != nil {
				log.Fatal(err)
			}
		case EventLeft:
			fmt.Println("LEFT")
			err = mouse.Turn(-1)
			if err != nil {
				log.Fatal(err)
			}
		case EventRight:
			fmt.Println("RIGHT")
			err = mouse.Turn(1)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
