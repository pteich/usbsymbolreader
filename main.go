package main

import (
	"context"
	"github.com/pteich/hid"
	"github.com/pteich/usbsymbolreader/scanner"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// create main context
	ctx, done := context.WithCancel(context.Background())

	// listen for system signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		select {
		case <-signalChannel:
			log.Print("shutdown signal received")
			done()
		}
	}()

	// main loop to enumerate USB devices and start reading from it
	for {

		log.Print("searching for USB devices...")
		devices := hid.Enumerate(0, 0)
		log.Printf("found %d devices", len(devices))

		for _, deviceInfo := range devices {
			log.Printf("found %s by %s . VendorID %d - ProductId %d", deviceInfo.Product, deviceInfo.Manufacturer, deviceInfo.VendorID, deviceInfo.ProductID)

			if deviceInfo.VendorID == 1504 || deviceInfo.Product == "Symbol Bar Code Scanner" {

				symbolScanner, err := scanner.New(deviceInfo)
				if err != nil {
					log.Print(err)
					break
				}

				log.Printf("connected to %s (Serial No. %s)", deviceInfo.Product, deviceInfo.Serial)

				codes := symbolScanner.ReadCodes(ctx)
				for code := range codes {
					// TODO safe to file
					log.Printf("scanned code: %s", code.String())
				}

				break
			}
		}

		select {
		case <-ctx.Done():
			log.Print("shutting down")
			return
		case <-time.After(1 * time.Second):
		}

	}
}
