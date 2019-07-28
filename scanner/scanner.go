package scanner

import (
	"context"
	"fmt"
	"github.com/pteich/hid"
	"github.com/pteich/usbsymbolreader/code"
)

// Scanner describes the symbol scanner device
type Scanner struct {
	device *hid.Device
}

// New takes a HID device info and returns a Scanner with the opened device
func New(deviceInfo hid.DeviceInfo) (*Scanner, error) {

	device, err := deviceInfo.Open()
	if err != nil {
		return nil, err
	}

	return &Scanner{
		device: device,
	}, nil
}

// ReadCodes starts a code read loop and returns a channel that recieves new codes
func (s *Scanner) ReadCodes(ctx context.Context) <-chan *code.Code {

	scanCtx, cancel := context.WithCancel(ctx)

	codeChan := make(chan *code.Code)
	go func() {
		defer close(codeChan)
		for {
			buf := make([]byte, 255)
			n, err := s.device.ReadTimeout(buf, 500)
			if err != nil {
				cancel()
				return
			}

			if n > 0 {
				fmt.Println(buf)
				scannedCode, err := code.New(buf)
				if err != nil {
					continue
				}

				codeChan <- scannedCode
			}

			select {
			case <-scanCtx.Done():
				s.device.Close()
				return
			default:
			}
		}
	}()

	return codeChan
}

// Device returns the real device
func (s *Scanner) Device() *hid.Device {
	return s.device
}
