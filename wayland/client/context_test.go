package client_test

import (
	"errors"
	"fmt"
	"github.com/MatthiasKunnen/go-wayland/wayland/client"
	"log"
	"time"
)

// Shows a dispatch loop that will block the goroutine.
// This approach has no risk of data races but the loop blocks the goroutine when no messages are
// received. This can be a valid approach if there are no more changes that need to be made after
// setting up and starting the loop.
func ExampleContext_Dispatch_single_routine() {
	display, err := client.Connect("")
	if err != nil {
		log.Fatalf("Error connecting to Wayland server: %v", err)
	}

	registry, err := display.GetRegistry()
	if err != nil {
		log.Fatalf("Error getting Wayland registry: %v", err)
	}

	var seat *client.Seat
	registry.SetGlobalHandler(func(e client.RegistryGlobalEvent) {
		switch e.Interface {
		case client.SeatInterfaceName:
			seat = client.NewSeat(display.Context())
			err := registry.Bind(e.Name, e.Interface, e.Version, seat)
			if err != nil {
				log.Fatalf("unable to bind %s interface: %v", client.SeatInterfaceName, err)
			}
		}
	})
	display.Roundtrip()
	display.Roundtrip()

	keyboard, err := seat.GetKeyboard()
	if err != nil {
		log.Printf("Error getting keyboard: %v", err)
	}
	log.Printf("Got keyboard: %v\n", keyboard)

	for {
		err := display.Context().Dispatch()
		if err != nil {
			log.Printf("Dispatch error: %v\n", err)
		}
	}
}

// Shows how the dispatch loop can be done in another goroutine.
// This prevents the goroutine from being blocked and allows making changes to wayland objects while
// the dispatch loop is blocking another goroutine.
func ExampleContext_Dispatch_multi_routine() {
	display, err := client.Connect("")
	if err != nil {
		log.Fatalf("Error connecting to Wayland server: %v", err)
	}

	registry, err := display.GetRegistry()
	if err != nil {
		log.Fatalf("Error getting Wayland registry: %v", err)
	}

	var seat *client.Seat
	registry.SetGlobalHandler(func(e client.RegistryGlobalEvent) {
		switch e.Interface {
		case client.SeatInterfaceName:
			seat = client.NewSeat(display.Context())
			err := registry.Bind(e.Name, e.Interface, e.Version, seat)
			if err != nil {
				log.Fatalf("unable to bind %s interface: %v", client.SeatInterfaceName, err)
			}
		}
	})
	display.Roundtrip()
	display.Roundtrip()

	go func() {
		err := display.Context().Dispatch()
		if err != nil {
			log.Printf("Dispatch error: %v\n", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	// The lock must be acquired to prevent data races.
	display.Context().Lock()
	defer display.Context().Unlock()
	keyboard, err := seat.GetKeyboard()
	if err != nil {
		log.Printf("Error getting keyboard: %v", err)
	}
	log.Printf("Got keyboard: %v\n", keyboard)

	err = errors.Join(keyboard.Release(), seat.Release(), display.Context().Close())
	if err != nil {
		fmt.Printf("Error cleaning up: %v\n", err)
	}
}
